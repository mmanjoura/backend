package activities

import (
	"net/http"
	"strconv"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// UpdateByID godoc
// @Summary Update a activity by ID
// @Description Update a activity by ID
// @Tags activities
// @Accept json
// @Produce json
// @Param id path int true "Activity ID"
// @Param activity body models.Activity true "Activity object"
// @Success 200 {string} string "Activity updated successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /activities/{id} [put]
func UpdateByID(c *gin.Context) {
	activityID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	var updatedActivity models.Activity

	if err := c.ShouldBindJSON(&updatedActivity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.Database.DB

	_, err = db.ExecContext(c, `
    UPDATE Activities
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
		updatedActivity.UserID,
		updatedActivity.Tag,
		updatedActivity.Title,
		updatedActivity.NumberOfReviews,
		updatedActivity.ReviewsComment,
		updatedActivity.Location,
		updatedActivity.Latitude,
		updatedActivity.Longitude,
		updatedActivity.MapUrl,
		updatedActivity.MinimumDuration,
		updatedActivity.GroupSize,
		updatedActivity.Overview,
		updatedActivity.CancellationPolicy,
		updatedActivity.WhatsIncluded,
		updatedActivity.Highlights,
		updatedActivity.AdditionalInformation,
		updatedActivity.ImportantInformation,
		updatedActivity.Price,
		updatedActivity.ActivityType,
		updatedActivity.Animation,
		time.Now(),
		activityID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while updating activity": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity updated successfully"})
}
