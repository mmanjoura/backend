package bookings

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
// @Tags bookings
// @Accept json
// @Produce json
// @Param booking body models.Booking true "FAQ object"
// @Success 200 {object} models.Booking "Successfully created a FAQ"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /bookings [post]
func Create(c *gin.Context) {
	var newBooking models.Booking

	productId := c.Query("product_id")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")
	numAdult := c.Query("num_adult")
	numChildren := c.Query("num_children")

	newBooking.ProductID = productId
	newBooking.StartDate = start_date
	newBooking.EndDate = end_date
	newBooking.NumAdult = numAdult
	newBooking.NumChildren = numChildren

	db := database.Database.DB

	result, err := db.ExecContext(c, `
			INSERT INTO bookings (
				product_id,
				start_date,
				end_date,
				product_type,
				adults,
				children,
				Created_At,
				Updated_At
			)
        VALUES ( ?, ?, ?, ?, ?, ?, ?, ?)`,
		newBooking.ProductID,
		newBooking.StartDate,
		newBooking.EndDate,
		"Activity",
		newBooking.NumAdult,
		newBooking.NumChildren,

		time.Now(),
		time.Now())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bookingId, _ := result.LastInsertId()
	newBooking.ID = int(bookingId)

	c.JSON(http.StatusOK, gin.H{"data": newBooking})
}
