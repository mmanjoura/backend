package contacts

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Get a contact by ID
// @Description Get a contact by ID
// @Tags contacts
// @Accept json
// @Produce json
// @Param id path int true "Contact ID"
// @Success 200 {object} contact "Contact found"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /contacts/{id} [get]
func GetByID(c *gin.Context) {
	contactID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}

	db := database.Database.DB

	var contact models.Contact
	err = db.QueryRowContext(c, `
	SELECT 	ID,
	name,
	email,
	subject,
	Created_At,
	Updated_At
	FROM contacts WHERE ID = ?`, contactID).
		Scan(&contact.ID,
			&contact.Name,
			&contact.Email,
			&contact.Subject,
			&contact.CreatedAt,
			&contact.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error while getting contact": "contact not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contact})
}
