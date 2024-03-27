package golfs

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage/backend/pkg/api/faqs"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/images"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/itineraries"
	"github.com/mmanjoura/niya-voyage/backend/pkg/common"
	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Get a golf by ID
// @Description Get a golf by ID
// @Tags golfs
// @Accept json
// @Produce json
// @Param id path int true "Golf ID"
// @Success 200 {object} models.Golf "Successfully retrieved a golf"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Golf not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /golfs/{id} [get]
func GetByID(c *gin.Context) {
	golfID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid golf ID"})
		return
	}

	db := database.Database.DB
	offset, limit := common.GetPaginationParams(c)
	categoryId := common.Activities

	var golf models.Golf
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
			golf_type,
			animation,
			Created_At,
			Updated_At
	FROM Golfs WHERE ID = ?`, golfID).
		Scan(&golf.ID,
			&golf.UserID,
			&golf.Tag,
			&golf.Title,
			&golf.NumberOfReviews,
			&golf.ReviewsComment,
			&golf.Location,
			&golf.Latitude,
			&golf.Longitude,
			&golf.MapUrl,
			&golf.MinimumDuration,
			&golf.GroupSize,
			&golf.Overview,
			&golf.CancellationPolicy,
			&golf.WhatsIncluded,
			&golf.Highlights,
			&golf.AdditionalInformation,
			&golf.ImportantInformation,
			&golf.Price,
			&golf.GolfType,
			&golf.Animation,
			&golf.CreatedAt,
			&golf.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error while getting Golf": "Golf not found"})
		return
	}

	if golf.ID != 0 {
		golfImages, err := images.RetrieveImages(c, db, limit, offset, golfID, categoryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error while getting GalleryImages": "Failed to get gallery images"})
			return
		}
		golfItineraries, err := itineraries.RetrieveItineraries(c, db, limit, offset, golf.ID, categoryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error while getting GolfItineraries": "Failed to get golf itineraries"})
			return
		}

		golfFaqs, err := faqs.RetrieveFaqs(c, db, limit, offset, golf.ID, categoryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error while getting GolfFaqs": "Failed to get golf faqs"})
			return
		}

		golf.Images = golfImages
		golf.Itineraries = golfItineraries
		golf.Faqs = golfFaqs
	}

	c.JSON(http.StatusOK, gin.H{"data": golf})
}
