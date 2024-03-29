package bookings

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/common"
	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// GetAll 			godoc
// @Summary 		Get all Bookings
// @Description 	Get all Bookings
// @Tags 			Bookings
// @Accept  		json
// @Produce 		json
// @Param 			offset query int false "Offset"
// @Param 			limit query int false "Limit"
// @Success 		200 {object} []Booking
// @Failure 		400 {string} string "Bad Request"
// @Failure 		500 {string} string "Internal Server Error"
// @Router 			/bookings [get]
func GetAll(c *gin.Context) {
	db := database.Database.DB
	offset, limit := common.GetPaginationParams(c)
	bookings, err := RetrieveBookings(c, db, limit, offset)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while retreiving Bookings": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookings})
}

func RetrieveBookings(c *gin.Context, db *sql.DB, limit, offset int) ([]models.Booking, error) {

	// Build WHERE clause. Each part of the clause is AND-ed together to further
	// restrict the results. Placeholders are added to "args" and are used
	// to avoid SQL injection.
	// Each filter field is optional.
	_, args := []string{"1 = 1"}, []interface{}{}

	bookings := []models.Booking{}

	rows, err := db.QueryContext(c, `SELECT ID,
			product_id,
			start_date,
			end_date,
			product_type,
			adults,
			children,
			Created_At,
			Updated_At,
			COUNT(*) OVER()
		FROM Bookings `,
		args...,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		booking, err := scanBooking(rows)

		if err != nil {
			return nil, err
		}

		if booking.ID != 0 {

			bookings = append(bookings, booking)
		}
	}
	defer rows.Close()

	return bookings, nil
}

func scanBooking(rows *sql.Rows) (models.Booking, error) {
	booking := models.Booking{}
	var n int

	err := rows.Scan(&booking.ID,
		&booking.ProductID,
		&booking.StartDate,
		&booking.EndDate,
		&booking.ProductType,
		&booking.NumAdult,
		&booking.NumChildren,
		(*time.Time)(&booking.CreatedAt),
		(*time.Time)(&booking.UpdatedAt),
		&n,
	)

	return booking, err
}
