package golfs

import (
	"net/http"
	"strconv"
	"time"

	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/models"

	"github.com/gin-gonic/gin"
)

// UpdateByID godoc
// @Summary Update a golf by ID
// @Description Update a golf by ID
// @Tags golfs
// @Accept json
// @Produce json
// @Param id path int true "Golf ID"
// @Param golf body models.Golf true "Golf object"
// @Success 200 {string} string "Golf updated successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /golfs/{id} [put]
func UpdateByID(c *gin.Context) {
	golfID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid golf ID"})
		return
	}

	var updatedGolf models.Golf

	if err := c.ShouldBindJSON(&updatedGolf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.Database.DB

	_, err = db.ExecContext(c, `
    UPDATE Golfs
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
			golf_type = ?,
			animation = ?,
			Updated_At = ?
		WHERE ID = ?`,
		updatedGolf.UserID,
		updatedGolf.Tag,
		updatedGolf.Title,
		updatedGolf.NumberOfReviews,
		updatedGolf.ReviewsComment,
		updatedGolf.Location,
		updatedGolf.Latitude,
		updatedGolf.Longitude,
		updatedGolf.MapUrl,
		updatedGolf.MinimumDuration,
		updatedGolf.GroupSize,
		updatedGolf.Overview,
		updatedGolf.CancellationPolicy,
		updatedGolf.WhatsIncluded,
		updatedGolf.Highlights,
		updatedGolf.AdditionalInformation,
		updatedGolf.ImportantInformation,
		updatedGolf.Price,
		updatedGolf.GolfType,
		updatedGolf.Animation,
		time.Now(),
		golfID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while updating golf": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Golf updated successfully"})
}
