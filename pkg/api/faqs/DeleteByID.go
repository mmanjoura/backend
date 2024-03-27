package faqs

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Delete a Faq by ID
// @Description Delete a Faq by ID
// @Tags faqs
// @Accept json
// @Produce json
// @Param id path int true "Faq ID"
// @Success 200 {string} string "Faq deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /faqs/{id} [delete]
func DeleteByID(c *gin.Context) {
	faqId, err := strconv.Atoi(c.Param("id"))
	db := database.Database.DB

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tour ID"})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM Faqs WHERE ID = ?`, faqId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting Faq": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Faq deleted successfully"})
}
