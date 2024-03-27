package contacts

import (
	"net/http"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create a new FAQ
// @Description Create a new FAQ
// @Tags contacts
// @Accept json
// @Produce json
// @Param contact body models.Contact true "FAQ object"
// @Success 200 {object} models.Contact "Successfully created a FAQ"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /contacts [post]
func Create(c *gin.Context) {
	var newContact models.Contact

	if err := c.ShouldBindJSON(&newContact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.Database.DB

	result, err := db.ExecContext(c, `
			INSERT INTO Contacts (
				name,
				email,
				subject,
				message,
				Created_At,
				Updated_At
			)
        VALUES ( ?, ?, ?, ?, ?, ?)`,

		newContact.Name,
		newContact.Email,
		newContact.Subject,
		newContact.Message,
		time.Now(),
		time.Now())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	newContact.ID = int(id)

	c.JSON(http.StatusOK, gin.H{"data": newContact})
}
