package contacts

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Delete a contact by ID
// @Description Delete a contact by ID
// @Tags contacts
// @Accept json
// @Produce json
// @Param id path int true "Contact ID"
// @Success 200 {string} string "Contact deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /contacts/{id} [delete]
func DeleteByID(c *gin.Context) {
	contactId, err := strconv.Atoi(c.Param("id"))
	db := database.Database.DB
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM contacts WHERE ID = ?`, contactId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting contact": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "contact deleted successfully"})
}
