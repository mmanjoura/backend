package tours

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/api/faqs"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/images"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/itineraries"
	"github.com/mmanjoura/niya-voyage/backend/pkg/common"
	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// Healthcheck godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /tours [get]
func Healthcheck(g *gin.Context) {
	g.JSON(http.StatusOK, "ok")
}

// GetAll 			godoc
// @Summary			Get all tours with pagination
// @Description 	Get a list of all tours with optional pagination
// @Tags 			tours
// @Security 		ApiKeyAuth
// @Produce 		json
// @Param 			offset query int false "Offset for pagination" default(0)
// @Param 			limit query int false "Limit for pagination" default(10)
// @Success 		200 {array} models.Tour "Successfully retrieved list of tours"
// @Router 			/tours [get]
func GetAll(c *gin.Context) {
	db := database.Database.DB
	offset, limit := common.GetPaginationParams(c)
	tours, err := retrieveTours(c, db, offset, limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while retreiving Tours": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tours})
}

func retrieveTours(c *gin.Context, db *sql.DB, offset, limit int) ([]models.Tour, error) {

	tours := []models.Tour{}
	categoryId := common.Tours

	rows, err := db.QueryContext(c, `SELECT ID,
			user_id,
			tag,
			title,
			number_of_reviews,
			reviews_comment,
			location,
			latitude,
			longitude,
			map_url,
			minimum_duration,
			group_size,
			overview,
			cancellation_policy,
			whats_included,
			highlights,
			additional_information,
			important_information,
			price,
			tour_type,
			animation,
			Created_At,
			Updated_At,
			COUNT(*) OVER()
		FROM Tours WHERE 1 = 1
		ORDER BY id desc `+database.FormatLimitOffset(limit, offset),
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tour, err := scanTour(rows)

		if err != nil {
			return nil, err
		}

		if tour.ID != 0 {
			tourImages, err := images.RetrieveImages(c, db, limit, offset, tour.ID, categoryId)
			if err != nil {
				return nil, err
			}

			tour.GalleryImages, err = images.RetrieveGalleryImages(c, db, limit, offset, tour.ID, categoryId)
			if err != nil {
				return nil, err
			}

			tour.SlideImages, err = images.RetrieveSlideImages(c, db, limit, offset, tour.ID, categoryId)
			if err != nil {
				return nil, err
			}

			tourItineraries, err := itineraries.RetrieveItineraries(c, db, limit, offset, tour.ID, categoryId)
			if err != nil {
				return nil, err
			}

			tourFaqs, err := faqs.RetrieveFaqs(c, db, limit, offset, tour.ID, categoryId)
			if err != nil {
				return nil, err
			}

			tour.Images = tourImages
			tour.GalleryImages = tour.GalleryImages
			tour.SlideImages = tour.SlideImages
			tour.Itineraries = tourItineraries
			tour.Faqs = tourFaqs
			tours = append(tours, tour)
		}
	}
	defer rows.Close()

	return tours, nil
}

func scanTour(rows *sql.Rows) (models.Tour, error) {
	tour := models.Tour{}
	var n int

	err := rows.Scan(&tour.ID,
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
		&tour.TourType,
		&tour.Animation,
		(*time.Time)(&tour.CreatedAt),
		(*time.Time)(&tour.UpdatedAt),
		&n,
	)

	return tour, err
}
