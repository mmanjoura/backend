package api

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"path"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/attribute"
	"github.com/mmanjoura/nomad/backend/category"

	// "github.com/mmanjoura/nomad/backend/categoryType"
	"github.com/mmanjoura/nomad/backend/coupon"
	"github.com/mmanjoura/nomad/backend/order"
	"github.com/mmanjoura/nomad/backend/product"
	"github.com/mmanjoura/nomad/backend/setting"
	"github.com/mmanjoura/nomad/backend/shipping"
	appUser "github.com/mmanjoura/nomad/backend/user"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// Generic HTTP metrics.
var (
	requestCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "app_http_request_count",
		Help: "Total number of requests by route",
	}, []string{"method", "path"})

	requestSeconds = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "app_http_request_seconds",
		Help: "Total amount of request time by route, in seconds",
	}, []string{"method", "path"})

	errorCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "app_http_error_count",
		Help: "Total number of errors by error code",
	}, []string{"code"})
)

// ShutdownTimeout is the time given for outstanding requests to finish before shutdown.
const ShutdownTimeout = 1 * time.Second

// Server represents an HTTP server. It is meant to wrap all HTTP functionality
// used by the application so that dependent packages (such as cmd/appd) do not
// need to reference the "net/http" package at all.
type Server struct {
	ln     net.Listener
	server *http.Server
	router *mux.Router

	sc *securecookie.SecureCookie

	// Bind address & domain for the server's listener.
	// If domain is specified, server is run on TLS using acme/autocert.
	Addr   string
	Domain string

	// Keys used for secure cookie encryption.
	HashKey  string
	BlockKey string

	// GitHub OAuth settings.
	GitHubClientID     string
	GitHubClientSecret string

	// Servics used by the various HTTP routes.
	AuthService appUser.AuthService
	// ShopService           user.ShopService
	// ShopMembershipService user.ShopMembershipService
	// EventService          user.EventService
	UserService      appUser.UserService
	AttributeService attribute.AttributeService
	CategoryService  category.CategoryService
	CouponService    coupon.CouponService
	ProductService   product.ProductService
	OrderService     order.OrderService
	SettingService   setting.SettingService
	ShippingService  shipping.ShippingService
	// CategoryTypeService categoryType.CategoryTypeService

}

// NewServer returns a new instance of Server.
func NewServer() *Server {
	// Create a new server that wraps the net/http server & add a gorilla router.

	s := &Server{
		server: &http.Server{},
		router: mux.NewRouter(),
	}

	// Report panics to external service.
	s.router.Use(reportPanic)

	// Our router is wrapped by another function handler to perform some
	// middleware-like tasks that cannot be performed by actual middleware.
	// This includes changing route paths for JSON endpoints & overridding methods.

	// handle CORS

	s.server.Handler = CORS(http.HandlerFunc(s.serveHTTP))

	// Setup error handling routes.
	s.router.NotFoundHandler = http.HandlerFunc(s.handleNotFound)

	// Handle embedded asset serving. This serves files embedded from http/assets.
	// s.router.PathPrefix("/assets/").
	// 	Handler(http.StripPrefix("/assets/", hashfs.FileServer(assets.FS)))

	// handler := NewHandler(s.UserService)

	// Setup endpoint to display deployed version.

	s.router.HandleFunc("/api/users", s.FindUsers).Methods("GET")
	s.router.HandleFunc("/api/users/{id}", s.FindUserByID).Methods("GET")
	s.router.HandleFunc("/api/users/{id}", s.DeleteUser).Methods("DELETE")
	s.router.HandleFunc("/api/users/{id}", s.UpdateUser).Methods("PUT")
	s.router.HandleFunc("/api/user", s.CreateUser).Methods("POST")

	s.router.HandleFunc("/api/attributes", s.FindAttributes).Methods("GET")

	s.router.HandleFunc("/api/categories", s.FindCategories).Methods("GET").Queries("searchJoin", "{searchJoin}", "limit", "{limit}", "page", "{page}", "parent", "{paent}", "search", "{search}")
	s.router.HandleFunc("/api/coupons", s.FindCoupons).Methods("GET")
	s.router.HandleFunc("/api/products", s.FindProducts).Methods("GET").Queries("searchJoin", "{searchJoin}", "limit", "{limit}", "page", "{page}", "parent", "{paent}", "search", "{search}")
	s.router.HandleFunc("/api/orders", s.FindOrders).Methods("GET")
	s.router.HandleFunc("/api/settings", s.FindSettings).Methods("GET")
	s.router.HandleFunc("/api/shippings", s.FindShippings).Methods("GET")
	s.router.HandleFunc("/api/types", s.FindCategoryTypes).Methods("GET")
	s.router.HandleFunc("/api/types/grocery", s.FindCategoryTypeByID).Methods("GET")
	// Setup a base router that excludes asset handling.
	router := s.router.PathPrefix("/").Subrouter()
	router.Use(s.authenticate)
	router.Use(loadFlash)
	router.Use(trackMetrics)

	// Handle authentication check within handler function for home page.
	//router.HandleFunc("/", s.handleIndex).Methods("GET")

	// Register unauthenticated routes.
	{
		r := s.router.PathPrefix("/").Subrouter()
		r.Use(s.requireNoAuth)
		s.registerAuthRoutes(r)
	}

	// Register authenticated routes.
	{
		r := router.PathPrefix("/").Subrouter()
		r.Use(s.requireAuth)
		r.HandleFunc("/settings", s.handleSettings).Methods("GET")
		//s.registerShopRoutes(r)
		//s.registerShopMembershipRoutes(r)
		//s.registerEventRoutes(r)

	}

	return s
}

// UseTLS returns true if the cert & key file are specified.
func (s *Server) UseTLS() bool {
	return s.Domain != ""
}

// Scheme returns the URL scheme for the server.
func (s *Server) Scheme() string {
	if s.UseTLS() {
		return "https"
	}
	return "http"
}

// Port returns the TCP port for the running server.
// This is useful in tests where we allocate a random port by using ":0".
func (s *Server) Port() int {
	if s.ln == nil {
		return 0
	}
	return s.ln.Addr().(*net.TCPAddr).Port
}

// URL returns the local base URL of the running server.
func (s *Server) URL() string {
	scheme, port := s.Scheme(), s.Port()

	// Use localhost unless a domain is specified.
	domain := "localhost"
	if s.Domain != "" {
		domain = s.Domain
	}

	// Return without port if using standard ports.
	if (scheme == "http" && port == 80) || (scheme == "https" && port == 443) {
		return fmt.Sprintf("%s://%s", s.Scheme(), domain)
	}
	return fmt.Sprintf("%s://%s:%d", s.Scheme(), domain, s.Port())
}

// Open validates the server options and begins listening on the bind address.
func (s *Server) Open() (err error) {
	// Initialize our secure cookie with our encryption keys.
	if err := s.openSecureCookie(); err != nil {
		return err
	}

	// Validate GitHub OAuth settings.
	if s.GitHubClientID == "" {
		return fmt.Errorf("github client id required")
	} else if s.GitHubClientSecret == "" {
		return fmt.Errorf("github client secret required")
	}

	// Open a listener on our bind address.
	if s.Domain != "" {
		s.ln = autocert.NewListener(s.Domain)
	} else {
		if s.ln, err = net.Listen("tcp", s.Addr); err != nil {
			return err
		}
	}

	// Begin serving requests on the listener. We use Serve() instead of
	// ListenAndServe() because it allows us to check for listen errors (such
	// as trying to use an already open port) synchronously.

	go s.server.Serve(s.ln)

	return nil
}

// openSecureCookie validates & decodes the block & hash key and initializes
// our secure cookie implementation.
func (s *Server) openSecureCookie() error {
	// Ensure hash & block key are set.
	if s.HashKey == "" {
		return fmt.Errorf("hash key required")
	} else if s.BlockKey == "" {
		return fmt.Errorf("block key required")
	}

	// Decode from hex to byte slices.
	hashKey, err := hex.DecodeString(s.HashKey)
	if err != nil {
		return fmt.Errorf("invalid hash key")
	}
	blockKey, err := hex.DecodeString(s.BlockKey)
	if err != nil {
		return fmt.Errorf("invalid block key")
	}

	// Initialize cookie management & encode our cookie data as JSON.
	s.sc = securecookie.New(hashKey, blockKey)
	s.sc.SetSerializer(securecookie.JSONEncoder{})

	return nil
}

// Close gracefully shuts down the server.
func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}

// OAuth2Config returns the GitHub OAuth2 configuration.
func (s *Server) OAuth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     s.GitHubClientID,
		ClientSecret: s.GitHubClientSecret,
		Scopes:       []string{},
		Endpoint:     github.Endpoint,
	}
}

func (s *Server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	// Override method for forms passing "_method" value.
	if r.Method == http.MethodPost {
		switch v := r.PostFormValue("_method"); v {
		case http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete:
			r.Method = v
		}
	}

	// Override content-type for certain extensions.
	// This allows us to easily cURL API endpoints with a ".json" or ".csv"
	// extension instead of having to explicitly set Content-type & Accept headers.
	// The extensions are removed so they don't appear in the routes.
	switch ext := path.Ext(r.URL.Path); ext {
	case ".json":
		r.Header.Set("Accept", "application/json")
		r.Header.Set("Content-type", "application/json")
		r.URL.Path = strings.TrimSuffix(r.URL.Path, ext)
	case ".csv":
		r.Header.Set("Accept", "text/csv")
		r.URL.Path = strings.TrimSuffix(r.URL.Path, ext)
	}

	// Delegate remaining HTTP handling to the gorilla router.
	s.router.ServeHTTP(w, r)
}

// authenticate is middleware for loading session data from a cookie or API key header.
func (s *Server) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Login via API key, if available.
		if v := r.Header.Get("Authorization"); strings.HasPrefix(v, "Bearer ") {
			apiKey := strings.TrimPrefix(v, "Bearer ")

			// Lookup user by API key. Display error if not found.
			// Otherwise set
			users, _, err := s.UserService.FindUsers(r.Context(), appUser.UserFilter{APIKey: &apiKey})
			if err != nil {
				Error(w, r, err)
				return
			} else if len(users) == 0 {
				Error(w, r, backend.Errorf(backend.EUNAUTHORIZED, "Invalid API key."))
				return
			}

			// Update request context to include authenticated user.
			r = r.WithContext(appUser.NewContextWithUser(r.Context(), users[0]))

			// Delegate to next HTTP handler.
			next.ServeHTTP(w, r)
			return
		}

		// Read session from secure cookie.
		session, _ := s.session(r)

		// Read user, if available. Ignore if fetching assets.
		if session.UserID != 0 {
			if user, err := s.UserService.FindUserByID(r.Context(), session.UserID); err != nil {
				log.Printf("cannot find session user: id=%d err=%s", session.UserID, err)
			} else {
				r = r.WithContext(appUser.NewContextWithUser(r.Context(), user))
			}
		}

		next.ServeHTTP(w, r)
	})
}

// requireNoAuth is middleware for requiring no authentication.
// This is used if a user goes to log in but is already logged in.
func (s *Server) requireNoAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If user is logged in, redirect to the home page.
		if userID := appUser.UserIDFromContext(r.Context()); userID != 0 {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		// Delegate to next HTTP handler.
		next.ServeHTTP(w, r)
	})
}

// requireAuth is middleware for requiring authentication. This is used by
// nearly every page except for the login & oauth pages.
func (s *Server) requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If user is logged in, delegate to next HTTP handler.
		if userID := appUser.UserIDFromContext(r.Context()); userID != 0 {
			next.ServeHTTP(w, r)
			return
		}

		// Otherwise save the current URL (without scheme/host).
		redirectURL := r.URL
		redirectURL.Scheme, redirectURL.Host = "", ""

		// Save the URL to the session and redirect to the log in page.
		// On successful login, the user will be redirected to their original location.
		session, _ := s.session(r)
		session.RedirectURL = redirectURL.String()
		if err := s.setSession(w, session); err != nil {
			log.Printf("http: cannot set session: %s", err)
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	})
}

// loadFlash is middleware for reading flash data from the cookie.
// Data is only loaded once and then immediately cleared... hence the name "flash".
func loadFlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read & clear flash from cookies.
		if cookie, _ := r.Cookie("flash"); cookie != nil {
			SetFlash(w, "")
			r = r.WithContext(appUser.NewContextWithFlash(r.Context(), cookie.Value))
		}

		// Delegate to next HTTP handler.
		next.ServeHTTP(w, r)
	})
}

// trackMetrics is middleware for tracking the request count and timing per route.
func trackMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtain path backend & start time of request.
		t := time.Now()
		tmpl := requestPathbackend(r)

		// Delegate to next handler in middleware chain.
		next.ServeHTTP(w, r)

		// Track total time unless it is the WebSocket endpoint for events.
		if tmpl != "" && tmpl != "/events" {
			requestCount.WithLabelValues(r.Method, tmpl).Inc()
			requestSeconds.WithLabelValues(r.Method, tmpl).Add(float64(time.Since(t).Seconds()))
		}
	})
}

// requestPathbackend returns the route path backend for r.
func requestPathbackend(r *http.Request) string {
	route := mux.CurrentRoute(r)
	if route == nil {
		return ""
	}
	tmpl, _ := route.GetPathTemplate()
	return tmpl
}

// reportPanic is middleware for catching panics and reporting them.
func reportPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				backend.ReportPanic(err)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// handleNotFound handles requests to routes that don't exist.
func (s *Server) handleNotFound(w http.ResponseWriter, r *http.Request) {
	// tmpl := html.Errorbackend{
	// 	StatusCode: http.StatusNotFound,
	// 	Header:     "Your page cannot be found.",
	// 	Message:    "Sorry, it looks like we can't find what you're looking for.",
	// }
	// tmpl.Render(r.Context(), w)
}

// handleIndex handles the "GET /" route. It displays a dashboard with the
// user's shops, recently updated membership values, & a chart.
func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	// If user is not logged in & application is built with a home page,
	// return the home page. Otherwise redirect to login.
	// if app.UserIDFromContext(r.Context()) == 0 {
	// 	if buf := assets.IndexHTML; len(buf) != 0 {
	// 		w.Write(buf)
	// 		return
	// 	}

	// 	http.Redirect(w, r, "/login", http.StatusFound)
	// 	return
	// }

	// var err error
	// var tmpl html.Indexbackend

	// Fetch all shops the user is a member of.
	// If user is not a member of any shops, redirect to shop list which
	// includes a description of how to start.
	// if tmpl.shops, _, err = s.ShopService.FindShops(r.Context(), app.ShopFilter{}); err != nil {
	// 	Error(w, r, err)
	// 	return
	// } else if len(tmpl.shops) == 0 {
	// 	http.Redirect(w, r, "/shops", http.StatusFound)
	// 	return
	// }

	// Fetch recently updated members.
	// if tmpl.Memberships, _, err = s.ShopMembershipService.FindShopMemberships(r.Context(), app.ShopMembershipFilter{
	// 	Limit:  20,
	// 	SortBy: "updated_at_desc",
	// }); err != nil {
	// 	Error(w, r, err)
	// 	return
	// }

	// Fetch historical average app values.
	// interval := time.Minute
	// end := time.Now().Truncate(interval).Add(interval)
	// start := end.Add(-1 * time.Hour)
	// if tmpl.AverageshopValueReport, err = s.ShopService.AverageShopValueReport(r.Context(), start, end, interval); err != nil {
	// 	Error(w, r, err)
	// 	return
	// }

	// Render the backend to the response.
	// tmpl.Render(r.Context(), w)
}

// handleSettings handles the "GET /settings" route.
func (s *Server) handleSettings(w http.ResponseWriter, r *http.Request) {
	// var tmpl html.Settingsbackend
	// tmpl.Render(r.Context(), w)
}

// handleVersion displays the deployed version.
func (s *Server) handleVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(backend.Version))
}

// handleVersion displays the deployed commit.
func (s *Server) handleCommit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(backend.Commit))
}

// session returns session data from the secure cookie.
func (s *Server) session(r *http.Request) (Session, error) {
	// Read session data from cookie.
	// If it returns an error then simply return an empty session.
	cookie, err := r.Cookie(SessionCookieName)
	if err != nil {
		return Session{}, nil
	}

	// Decode session data into a Session object & return.
	var session Session
	if err := s.UnmarshalSession(cookie.Value, &session); err != nil {
		return Session{}, err
	}
	return session, nil
}

// setSession creates a secure cookie with session data.
func (s *Server) setSession(w http.ResponseWriter, session Session) error {
	// Encode session data to JSON.
	buf, err := s.MarshalSession(session)
	if err != nil {
		return err
	}

	// Write cookie to HTTP response.
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    buf,
		Path:     "/",
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		Secure:   s.UseTLS(),
		HttpOnly: true,
	})
	return nil
}

// MarshalSession encodes session data to string.
// This is exported to allow the unit tests to generate fake sessions.
func (s *Server) MarshalSession(session Session) (string, error) {
	return s.sc.Encode(SessionCookieName, session)
}

// UnmarshalSession decodes session data into a Session object.
// This is exported to allow the unit tests to generate fake sessions.
func (s *Server) UnmarshalSession(data string, session *Session) error {
	return s.sc.Decode(SessionCookieName, data, &session)
}

// ListenAndServeTLSRedirect runs an HTTP server on port 80 to redirect users
// to the TLS-enabled port 443 server.
func ListenAndServeTLSRedirect(domain string) error {
	return http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://"+domain, http.StatusFound)
	}))
}

// ListenAndServeDebug runs an HTTP server with /debug endpoints (e.g. pprof, vars).
func ListenAndServeDebug() error {
	h := http.NewServeMux()
	h.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(":6060", h)
}

// Error prints & optionally logs an error message.
func Error(w http.ResponseWriter, r *http.Request, err error) {
	// Extract error code & message.
	code, message := backend.ErrorCode(err), backend.ErrorMessage(err)

	// Track metrics by code.
	errorCount.WithLabelValues(code).Inc()

	// Log & report internal errors.
	if code == backend.EINTERNAL {
		backend.ReportError(r.Context(), err, r)
		LogError(r, err)
	}

	// Print user message to response based on reqeust accept header.
	switch r.Header.Get("Accept") {
	case "application/json":
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(ErrorStatusCode(code))
		json.NewEncoder(w).Encode(&ErrorResponse{Error: message})

	default:
		w.WriteHeader(ErrorStatusCode(code))
		// tmpl := html.Errorbackend{
		// 	StatusCode: ErrorStatusCode(code),
		// 	Header:     "An error has occurred.",
		// 	Message:    message,
		// }
		// tmpl.Render(r.Context(), w)
	}
}

// ErrorStatusCode returns the associated HTTP status code for a app error code.
func ErrorStatusCode(code string) int {
	if v, ok := codes[code]; ok {
		return v
	}
	return http.StatusInternalServerError
}

// FromErrorStatusCode returns the associated app code for an HTTP status code.
func FromErrorStatusCode(code int) string {
	for k, v := range codes {
		if v == code {
			return k
		}
	}
	return backend.EINTERNAL
}

// lookup of application error codes to HTTP status codes.
var codes = map[string]int{
	backend.ECONFLICT:       http.StatusConflict,
	backend.EINVALID:        http.StatusBadRequest,
	backend.ENOTFOUND:       http.StatusNotFound,
	backend.ENOTIMPLEMENTED: http.StatusNotImplemented,
	backend.EUNAUTHORIZED:   http.StatusUnauthorized,
	backend.EINTERNAL:       http.StatusInternalServerError,
}

// LogError logs an error with the HTTP route information.
func LogError(r *http.Request, err error) {
	log.Printf("[http] error: %s %s: %s", r.Method, r.URL.Path, err)
}

// ErrorResponse represents a JSON structure for error output.
type ErrorResponse struct {
	Error string `json:"error"`
}

// SessionCookieName is the name of the cookie used to store the session.
const SessionCookieName = "session"

// Session represents session data stored in a secure cookie.
type Session struct {
	UserID      int    `json:"userID"`
	RedirectURL string `json:"redirectURL"`
	State       string `json:"state"`
}

// SetFlash sets the flash cookie for the next request to read.
func SetFlash(w http.ResponseWriter, s string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "flash",
		Value:    s,
		Path:     "/",
		HttpOnly: true,
	})
}

func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST")

			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
			return
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
