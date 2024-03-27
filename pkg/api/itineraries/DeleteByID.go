package itineraries

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Delete a itinerary by ID
// @Description Delete a itinerary by ID
// @Tags itinerarys
// @Accept json
// @Produce json
// @Param id path int true "Itinerary ID"
// @Success 200 {string} string "Itinerary deleted successfully"
// @Failure 400 {string} string "Bad Request"
func DeleteByID(c *gin.Context) {
	itineraryID, err := strconv.Atoi(c.Param("id"))
	db := database.Database.DB

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tour ID"})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM Itineraries WHERE ID = ?`, itineraryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting Itinerary": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Itinerary deleted successfully"})
}
