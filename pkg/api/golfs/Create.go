package golfs

import (
	"net/http"
	"time"

	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create a new golf
// @Description Create a new golf
// @Tags golfs
// @Accept json
// @Produce json
// @Param golf body models.Golf true "Golf object"
// @Success 200 {object} models.Golf "Successfully created a new golf"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /golfs [post]

func Create(c *gin.Context) {
	var newGolf models.CreateGolf

	if err := c.ShouldBindJSON(&newGolf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add additional validations if needed

	db := database.Database.DB

	result, err := db.ExecContext(c, `
		INSERT INTO Golfs (
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
			golf_type,
			animation,
			Created_At,
			Updated_At)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,

		newGolf.UserID,
		newGolf.Tag,
		newGolf.Title,
		newGolf.NumberOfReviews,
		newGolf.ReviewsComment,
		newGolf.Location,
		newGolf.Latitude,
		newGolf.Longitude,
		newGolf.MapUrl,
		newGolf.MinimumDuration,
		newGolf.GroupSize,
		newGolf.Overview,
		newGolf.WhatsIncluded,
		newGolf.Highlights,
		newGolf.CancellationPolicy,
		newGolf.AdditionalInformation,
		newGolf.ImportantInformation,
		newGolf.Price,
		newGolf.GolfType,
		"100",
		time.Now(),
		time.Now())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	newGolf.ID = int(id)

	c.JSON(http.StatusOK, gin.H{"data": newGolf})
}
