package bookings

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Get a booking by ID
// @Description Get a booking by ID
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} booking "Booking found"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /bookings/{id} [get]
func GetByID(c *gin.Context) {
	bookingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	db := database.Database.DB

	var booking models.Booking
	err = db.QueryRowContext(c, `
	SELECT 	ID,
	product_id,
	start_date,
	end_date,
	adults,
	children,
	product_type,
	Created_At,
	Updated_At
	FROM bookings WHERE ID = ?`, bookingID).
		Scan(&booking.ID,
			&booking.ProductID,
			&booking.StartDate,
			&booking.EndDate,
			&booking.ProductType,
			&booking.NumAdult,
			&booking.NumChildren,
			&booking.CreatedAt,
			&booking.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error while getting booking": "booking not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booking})
}
