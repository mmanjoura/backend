package middleware

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func Security() gin.HandlerFunc {
	return secure.New(secure.Config{
		AllowedHosts:          []string{"niyavoyage.com", "ssl.niyavoyage.com", "https://niya-voyage-backend-app-d4a23urhsq-uc.a.run.app/api/v1"},
		SSLRedirect:           true,
		SSLHost:               "https://niya-voyage-backend-app-d4a23urhsq-uc.a.run.app/api/v1",
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen:              true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	})
}
