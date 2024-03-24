package activities

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/api/faqs"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/api/images"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/api/itineraries"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/common"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/models"

	"github.com/gin-gonic/gin"
)

// GetAll 			godoc
// @Summary			Get all activities with pagination
// @Description 	Get a list of all activities with optional pagination
// @Tags 			activities
// @Security 		ApiKeyAuth
// @Produce 		json
// @Param 			offset query int false "Offset for pagination" default(0)
// @Param 			limit query int false "Limit for pagination" default(10)
// @Success 		200 {array} models.Activity "Successfully retrieved list of activities"
// @Router 			/activities [get]
func GetAll(c *gin.Context) {
	db := database.Database.DB
	offset, limit := common.GetPaginationParams(c)
	activities, err := retrieveActivities(c, db, offset, limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while retreiving Activitys": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": activities})
}

func retrieveActivities(c *gin.Context, db *sql.DB, offset, limit int) ([]models.Activity, error) {

	activities := []models.Activity{}
	categoryId := common.Activities

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
		FROM Activities WHERE 1 = 1
		ORDER BY id desc `+database.FormatLimitOffset(limit, offset),
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		activity, err := scanActivity(rows)

		if err != nil {
			return nil, err
		}

		if activity.ID != 0 {
			activityImages, err := images.RetrieveImages(c, db, limit, offset, activity.ID, categoryId)
			if err != nil {
				return nil, err
			}

			activityItineraries, err := itineraries.RetrieveItineraries(c, db, limit, offset, activity.ID, categoryId)
			if err != nil {
				return nil, err
			}

			activityFaqs, err := faqs.RetrieveFaqs(c, db, limit, offset, activity.ID, categoryId)
			if err != nil {
				return nil, err
			}

			activity.Images = activityImages
			activity.Itineraries = activityItineraries
			activity.Faqs = activityFaqs
			activities = append(activities, activity)
		}
	}
	defer rows.Close()

	return activities, nil
}

func scanActivity(rows *sql.Rows) (models.Activity, error) {
	activity := models.Activity{}
	var n int

	err := rows.Scan(&activity.ID,
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
		(*time.Time)(&activity.CreatedAt),
		(*time.Time)(&activity.UpdatedAt),
		&n,
	)

	return activity, err
}
