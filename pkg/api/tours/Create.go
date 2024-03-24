package tours

import (
	"net/http"
	"time"

	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create a new tour
// @Description Create a new tour
// @Tags tours
// @Accept json
// @Produce json
// @Param tour body models.Tour true "Tour object"
// @Success 200 {object} models.Tour "Successfully created a new tour"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tours [post]

func Create(c *gin.Context) {
	var newTour models.CreateTour

	if err := c.ShouldBindJSON(&newTour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add additional validations if needed

	db := database.Database.DB

	result, err := db.ExecContext(c, `
		INSERT INTO Tours (
			user_id,
			tag,
			title,
			number_of_reviews,
			reviews_comment,
			location,
			latitude,
			longitude,
			map_url ,
			minimum_duration,
			group_size,
			overview,
			whats_included,
			highlights,
			cancellation_policy,
			additional_information,
			important_information,
			price,
			tour_type,
			animation,
			Created_At,
			Updated_At)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		newTour.UserID,
		newTour.Tag,
		newTour.Title,
		newTour.NumberOfReviews,
		newTour.ReviewsComment,
		newTour.Location,
		newTour.Latitude,
		newTour.Longitude,
		newTour.MapUrl,
		newTour.MinimumDuration,
		newTour.GroupSize,
		newTour.Overview,
		newTour.WhatsIncluded,
		newTour.Highlights,
		newTour.CancellationPolicy,
		newTour.AdditionalInformation,
		newTour.ImportantInformation,
		newTour.Price,
		newTour.TourType,
		"100",
		time.Now(),
		time.Now())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	newTour.ID = int(id)

	c.JSON(http.StatusOK, gin.H{"data": newTour})
}
