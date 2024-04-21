package tours

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
// @Summary Get a tour by ID
// @Description Get a tour by ID
// @Tags tours
// @Accept json
// @Produce json
// @Param id path int true "Tour ID"
// @Success 200 {object} models.Tour "Successfully retrieved a tour"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Tour not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tours/{id} [get]
func GetByID(c *gin.Context) {
	tourID, err := strconv.Atoi(c.Param("id"))
	categoryId := common.Tours
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tour ID"})
		return
	}

	db := database.Database.DB
	offset, limit := common.GetPaginationParams(c)

	var tour models.Tour
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
	FROM Tours WHERE ID = ?`, tourID).
		Scan(&tour.ID,
			&tour.UserID,
			&tour.Tag,
			&tour.Title,
			&tour.NumberOfReviews,
			&tour.ReviewsComment,
			&tour.Location,
			&tour.Latitude,
			&tour.Longitude,
			&tour.MapUrl,
			&tour.MinimumDuration,
			&tour.GroupSize,
			&tour.Overview,
			&tour.CancellationPolicy,
			&tour.WhatsIncluded,
			&tour.Highlights,
			&tour.AdditionalInformation,
			&tour.ImportantInformation,
			&tour.Price,
			&tour.ActivityType,
			&tour.Animation,
			&tour.CreatedAt,
			&tour.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error while getting Tour": "Tour not found"})
		return
	}

	if tour.ID != 0 {
		tourImages, err := images.RetrieveImages(c, db, limit, offset, tour.ID, categoryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error while getting GalleryImages": "Failed to get gallery images"})
			return
		}

		tourGalleryImages, err := images.RetrieveGalleryImages(c, db, limit, offset, tour.ID, categoryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error while getting GalleryImages": "Failed to get gallery images"})
			return
		}

		tourSlideImages, err := images.RetrieveSlideImages(c, db, limit, offset, tour.ID, categoryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error while getting SlideImages": "Failed to get slide images"})
			return
		}

		tourItineraries, err := itineraries.RetrieveItineraries(c, db, limit, offset, tour.ID, categoryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error while getting TourItineraries": "Failed to get tour itineraries"})
			return
		}

		tourFaqs, err := faqs.RetrieveFaqs(c, db, limit, offset, tour.ID, categoryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error while getting TourFaqs": "Failed to get tour faqs"})
			return
		}

		tour.Images = tourImages
		tour.SlideImages = tourSlideImages
		tour.GalleryImages = tourGalleryImages
		tour.Itineraries = tourItineraries
		tour.Faqs = tourFaqs
	}

	c.JSON(http.StatusOK, gin.H{"data": tour})
}
