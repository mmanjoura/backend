package middleware

import (
	"net/http"
	"time"

	auth "github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/auth"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("Authorization")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "cannot Get Authorization token"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return auth.JwtKey, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Cannot parse token"})
			c.Abort()
			return
		}

		claims := jwt.MapClaims{}
		var ok bool
		if claims, ok = token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := claims["exp"].(float64)
			if exp < float64(time.Now().Unix()) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
				c.Abort()
				return
			}

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token when mapping Claims"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		user, err := getUser(c, claims["iss"].(string))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Error when getting user from token"})
			c.Abort()
			return
		}

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not found"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func getUser(c *gin.Context, userEmail string) (models.User, error) {
	db := database.Database.DB

	user := models.User{}

	err := db.QueryRowContext(c, `
	SELECT id,
       firstname,
       lastname,
       email,
       password,
	   isAdmin,
       Updated_At,
       Created_At
  FROM Users WHERE email = ?`, userEmail).
		Scan(&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.IsAdmin,
			&user.UpdatedAt,
			&user.CreatedAt,
		)

	if err != nil {
		return user, err
	}

	return user, nil
}
