package bookings

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Delete a booking by ID
// @Description Delete a booking by ID
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {string} string "Booking deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /bookings/{id} [delete]
func DeleteByID(c *gin.Context) {
	bookingId, err := strconv.Atoi(c.Param("id"))
	db := database.Database.DB
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM bookings WHERE ID = ?`, bookingId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting booking": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "booking deleted successfully"})
}
