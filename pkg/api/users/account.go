package users

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/common"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/models"

	"github.com/gin-gonic/gin"
)

// UserAccount godoc
// @Summary   User account
// @Tags      user
// @Accept    json
// @Produce   json
// @Success   200  {object}  auth.AuthDetails
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  UserAuth
// @Router    /users/account [get]
func Account(c *gin.Context) {
	offset, limit := common.GetPaginationParams(c)
	// userId, err := strconv.Atoi(c.Param("id"))
	// Retrieve the value of the "jwt" cookie
	jwtCookie, err := c.Cookie("Authorization")
	if err != nil {
		// Handle the case when the cookie is not found or there's an error
		c.JSON(400, gin.H{"error": "Cookie not found"})
		return
	}

	token, err := jwt.ParseWithClaims(jwtCookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("k1U6pO+9qZteWy+yE52Z56qSBqmJ1orl27r/28AfkIA="), nil
	})

	if err != nil {
		// Handle parsing errors
		c.JSON(400, gin.H{"error": "Invalid token: " + err.Error()})
		return
	}

	if !token.Valid {
		// Handle invalid token
		c.JSON(400, gin.H{"error": "Invalid token"})
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	account, err := retrieveUserAccount(c, offset, limit, claims.Issuer)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while retreiving Accounts": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": account})
}

func retrieveUserAccount(c *gin.Context, offset, limit int, userEmail string) (models.User, error) {
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
