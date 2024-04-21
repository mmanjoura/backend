package tours

import (
	"net/http"
	"strconv"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// UpdateByID godoc
// @Summary Update a tour by ID
// @Description Update a tour by ID
// @Tags tours
// @Accept json
// @Produce json
// @Param id path int true "Tour ID"
// @Param tour body models.Tour true "Tour object"
// @Success 200 {string} string "Tour updated successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tours/{id} [put]
func UpdateByID(c *gin.Context) {
	tourID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tour ID"})
		return
	}

	var updatedTour models.Tour

	if err := c.ShouldBindJSON(&updatedTour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.Database.DB

	_, err = db.ExecContext(c, `
    UPDATE Tours
		SET user_id = ?,
			tag = ?,
			title = ?,
			number_of_reviews = ?,
			reviews_comment = ?,
			location = ?,
			latitude = ?,
			longitude = ?,
			map_url = ?,
			minimum_duration = ?,
			group_size = ?,
			overview = ?,
			cancellation_policy = ?,
			whats_included = ?,
			highlights = ?,
			additional_information = ?,
			important_information = ?,
			price = ?,
			activity_type = ?,
			animation = ?,
			Updated_At = ?
		WHERE ID = ?`,
		updatedTour.UserID,
		updatedTour.Tag,
		updatedTour.Title,
		updatedTour.NumberOfReviews,
		updatedTour.ReviewsComment,
		updatedTour.Location,
		updatedTour.Latitude,
		updatedTour.Longitude,
		updatedTour.MapUrl,
		updatedTour.MinimumDuration,
		updatedTour.GroupSize,
		updatedTour.Overview,
		updatedTour.CancellationPolicy,
		updatedTour.WhatsIncluded,
		updatedTour.Highlights,
		updatedTour.AdditionalInformation,
		updatedTour.ImportantInformation,
		updatedTour.Price,
		updatedTour.ActivityType,
		updatedTour.Animation,
		time.Now(),
		tourID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while updating tour": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tour updated successfully"})
}
