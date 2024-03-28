package contacts

import (
	"net/http"
	"strconv"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// UpdateByID godoc
// @Summary Update a contact by ID
// @Description Update a contact by ID
// @Tags contacts
// @Accept json
// @Produce json
// @Param id path int true "Contact ID"
// @Param contact body contact true "Contact object"
// @Success 200 {string} string "Tour updated successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /contacts/{id} [put]
func UpdateByID(c *gin.Context) {
	contactID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}

	var updatedContact models.Contact

	if err := c.ShouldBindJSON(&updatedContact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.Database.DB
	defer db.Close()

	_, err = db.ExecContext(c, `
								UPDATE Contacts
								SET name = ?,
									email = ?,
									subject = ?,
									Updated_At = ?
							WHERE id = ?`,
		updatedContact.Name, updatedContact.Email,
		updatedContact.Subject,
		time.Now(), contactID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while updating contact": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tour updated successfully"})
}
