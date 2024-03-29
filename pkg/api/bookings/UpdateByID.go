package bookings

import (
	"net/http"
	"strconv"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// UpdateByID godoc
// @Summary Update a booking by ID
// @Description Update a booking by ID
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Param booking body booking true "Booking object"
// @Success 200 {string} string "Tour updated successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /bookings/{id} [put]
func UpdateByID(c *gin.Context) {
	bookingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	var updatedBooking models.Booking

	if err := c.ShouldBindJSON(&updatedBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.Database.DB
	defer db.Close()

	_, err = db.ExecContext(c, `
								UPDATE Bookings
								SET   product_id = ?,
								start_date = ?,
								end_date = ?,
								product_type = ?,
								adults = ?,
								children = ?,								
								Updated_At = ?
							WHERE id = ?`,
		updatedBooking.ProductID,
		updatedBooking.StartDate,
		updatedBooking.EndDate,
		updatedBooking.ProductType,
		updatedBooking.NumAdult,
		updatedBooking.NumChildren,

		time.Now(), bookingID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while updating booking": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tour updated successfully"})
}
