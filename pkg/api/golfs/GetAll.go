package golfs

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

// GetAll 			godoc
// @Summary			Get all golfs with pagination
// @Description 	Get a list of all golfs with optional pagination
// @Tags 			golfs
// @Security 		ApiKeyAuth
// @Produce 		json
// @Param 			offset query int false "Offset for pagination" default(0)
// @Param 			limit query int false "Limit for pagination" default(10)
// @Success 		200 {array} models.Golf "Successfully retrieved list of golfs"
// @Router 			/golfs [get]
func GetAll(c *gin.Context) {
	db := database.Database.DB
	offset, limit := common.GetPaginationParams(c)
	golfs, err := retrieveActivities(c, db, offset, limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while retreiving Golfs": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": golfs})
}

func retrieveActivities(c *gin.Context, db *sql.DB, offset, limit int) ([]models.Golf, error) {

	golfs := []models.Golf{}
	categoryId := common.Golfs

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
			activity_type,
			animation,
			Created_At,
			Updated_At,
			COUNT(*) OVER()
		FROM Golfs WHERE 1 = 1
		ORDER BY id desc `+database.FormatLimitOffset(limit, offset),
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		golf, err := scanGolf(rows)

		if err != nil {
			return nil, err
		}

		if golf.ID != 0 {
			golfImages, err := images.RetrieveImages(c, db, limit, offset, golf.ID, categoryId)
			if err != nil {
				return nil, err
			}
			golf.GalleryImages, err = images.RetrieveGalleryImages(c, db, limit, offset, golf.ID, categoryId)
			if err != nil {
				return nil, err
			}

			golf.SlideImages, err = images.RetrieveSlideImages(c, db, limit, offset, golf.ID, categoryId)
			if err != nil {
				return nil, err
			}

			golfItineraries, err := itineraries.RetrieveItineraries(c, db, limit, offset, golf.ID, categoryId)
			if err != nil {
				return nil, err
			}

			golfFaqs, err := faqs.RetrieveFaqs(c, db, limit, offset, golf.ID, categoryId)
			if err != nil {
				return nil, err
			}

			golf.Images = golfImages
			golf.Itineraries = golfItineraries
			golf.Faqs = golfFaqs
			golfs = append(golfs, golf)
		}
	}
	defer rows.Close()

	return golfs, nil
}

func scanGolf(rows *sql.Rows) (models.Golf, error) {
	golf := models.Golf{}
	var n int

	err := rows.Scan(&golf.ID,
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
		&golf.ActivityType,
		&golf.Animation,
		(*time.Time)(&golf.CreatedAt),
		(*time.Time)(&golf.UpdatedAt),
		&n,
	)

	return golf, err
}
