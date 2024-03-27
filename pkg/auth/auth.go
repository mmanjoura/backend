package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Claims struct to be encoded to JWT
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// var config = database.Database.Config
// var JwtKey = []byte(config["JWT-API-KEY"])
var JwtKey = []byte("k1U6pO+9qZteWy+yE52Z56qSBqmJ1orl27r/28AfkIA=")

// @BasePath /api/v1

// LoginHandler godoc
// @Summary Authenticate a user
// @Schemes
// @Description Authenticates a user using username and password, returns a JWT token if successful
// @Tags user
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param   user     body    models.LoginUser     true        "User login object"
// @Success 200 {string} string "JWT Token"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	var incomingUser models.LoginUser
	var dbUser models.User
	db := database.Database.DB

	// Get JSON body
	if err := c.ShouldBindJSON(&incomingUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	// Fetch the user from the database
	err := db.QueryRow("SELECT id, firstname, lastname, email, password, isAdmin, Created_At, Updated_At FROM users WHERE email = ?", incomingUser.Email).
		Scan(&dbUser.ID, &dbUser.FirstName, &dbUser.LastName, &dbUser.Email, &dbUser.Password, &dbUser.IsAdmin, &dbUser.CreatedAt, &dbUser.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(incomingUser.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT token
	token, err := GenerateToken(dbUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}
	c.SetSameSite(http.SameSiteNoneMode)
	// Set a cookie named "Authorization" with the provided token value
	// Replace "yourdomain.com" with the appropriate domain for your Cloud Run service
	c.SetCookie("Authorization", token, int(time.Hour*24), "/", "", true, false)

	c.JSON(http.StatusOK, gin.H{"user": dbUser})
}

// RegisterHandler godoc
// @Summary Register a new user
// @Schemes http
// @Description Registers a new user with the given username and password
// @Tags user
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param   user     body    models.LoginUser     true        "User registration object"
// @Success 200 {string} string	"Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /register [post]
func RegisterHandler(c *gin.Context) {
	var user models.User
	db := database.Database.DB

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	// Create new user
	newUser := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  hashedPassword,
		IsAdmin:   0,
		Email:     user.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Execute the SQL query to insert a new user
	_, err = db.Exec("INSERT INTO users (firstname, lastname, email, password, Created_At, Updated_At) VALUES (?, ?, ?, ?, ?, ?)",
		newUser.FirstName, newUser.LastName, newUser.Email, newUser.Password, newUser.CreatedAt, newUser.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateToken(email string) (string, error) {
	// The expiration time after which the token will be invalid.
	expirationTime := time.Now().Add(12 * time.Hour).Unix()

	// Create the JWT claims, which includes the email and expiration time
	claims := &jwt.StandardClaims{
		// In JWT, the expiry time is expressed as unix milliseconds
		ExpiresAt: expirationTime,
		Issuer:    email,
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// logout godoc
// @Summary   Logout
// @Tags      user
// @Accept    json
// @Produce   json
// @Success   200  {string}  success
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  UserAuth
// @Router    /logout [get]
func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", 0, "/", "localhost", false, true)
	c.JSON(http.StatusOK, "success")
}

func GenerateRandomKey() string {
	key := make([]byte, 32) // generate a 256 bit key
	_, err := rand.Read(key)
	if err != nil {
		panic("Failed to generate random key: " + err.Error())
	}

	return base64.StdEncoding.EncodeToString(key)
}
