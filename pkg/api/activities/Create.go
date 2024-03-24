package activities

import (
	"net/http"
	"time"

	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create a new activity
// @Description Create a new activity
// @Tags activities
// @Accept json
// @Produce json
// @Param activity body models.Activity true "Activity object"
// @Success 200 {object} models.Activity "Successfully created a new activity"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /activities [post]

func Create(c *gin.Context) {
	var newActivity models.CreateActivity

	if err := c.ShouldBindJSON(&newActivity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add additional validations if needed

	db := database.Database.DB

	result, err := db.ExecContext(c, `
		INSERT INTO Activities (
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
			activity_type,
			animation,
			Created_At,
			Updated_At)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,

		newActivity.UserID,
		newActivity.Tag,
		newActivity.Title,
		newActivity.NumberOfReviews,
		newActivity.ReviewsComment,
		newActivity.Location,
		newActivity.Latitude,
		newActivity.Longitude,
		newActivity.MapUrl,
		newActivity.MinimumDuration,
		newActivity.GroupSize,
		newActivity.Overview,
		newActivity.WhatsIncluded,
		newActivity.Highlights,
		newActivity.CancellationPolicy,
		newActivity.AdditionalInformation,
		newActivity.ImportantInformation,
		newActivity.Price,
		newActivity.ActivityType,
		"100",
		time.Now(),
		time.Now())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	newActivity.ID = int(id)

	c.JSON(http.StatusOK, gin.H{"data": newActivity})
}
