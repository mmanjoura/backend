package activities

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/api/faqs"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/api/images"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/api/itineraries"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/common"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Get a activity by ID
// @Description Get a activity by ID
// @Tags activities
// @Accept json
// @Produce json
// @Param id path int true "Activity ID"
// @Success 200 {object} models.Activity "Successfully retrieved a activity"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Activity not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /activities/{id} [get]
func GetByID(c *gin.Context) {
	activityID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	db := database.Database.DB
	offset, limit := common.GetPaginationParams(c)
	categoryId := common.Activities

	var activity models.Activity
	err = db.QueryRowContext(c, `
	SELECT 	ID,
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
			cancellation_policy,
			whats_included,
			highlights,
			additional_information,
			important_information,
			price,
			activity_type,
			animation,
			Created_At,
			Updated_At
	FROM Activities WHERE ID = ?`, activityID).
		Scan(&activity.ID,
			&activity.UserID,
			&activity.Tag,
			&activity.Title,
			&activity.NumberOfReviews,
			&activity.ReviewsComment,
			&activity.Location,
			&activity.Latitude,
			&activity.Longitude,
			&activity.MapUrl,
			&activity.MinimumDuration,
			&activity.GroupSize,
			&activity.Overview,
			&activity.CancellationPolicy,
			&activity.WhatsIncluded,
			&activity.Highlights,
			&activity.AdditionalInformation,
			&activity.ImportantInformation,
			&activity.Price,
			&activity.ActivityType,
			&activity.Animation,
			&activity.CreatedAt,
			&activity.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error while getting Activity": "Activity not found"})
		return
	}

	if activity.ID != 0 {
		activityImages, err := images.RetrieveImages(c, db, limit, offset, activityID, categoryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error while getting GalleryImages": "Failed to get gallery images"})
			return
		}
		activityItineraries, err := itineraries.RetrieveItineraries(c, db, limit, offset, activity.ID, categoryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error while getting ActivityItineraries": "Failed to get activity itineraries"})
			return
		}

		activityFaqs, err := faqs.RetrieveFaqs(c, db, limit, offset, activity.ID, categoryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error while getting ActivityFaqs": "Failed to get activity faqs"})
			return
		}

		activity.Images = activityImages
		activity.Itineraries = activityItineraries
		activity.Faqs = activityFaqs
	}

	c.JSON(http.StatusOK, gin.H{"data": activity})
}
