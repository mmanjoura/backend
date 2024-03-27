package itineraries

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/common"
	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// GetAll 			godoc
// @Summary Get all Itineraries
// @Description Get all Itineraries
// @Tags itinerarys
// @Accept json
// @Produce json
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Success 200 {array} models.Itinerary "Successfully retrieved list of Itineraries"
// @Router /itineraries [get]

func GetAll(c *gin.Context) {
	db := database.Database.DB
	limit, offset := common.GetPaginationParams(c)
	itineraries, err := RetrieveItineraries(c, db, limit, offset, 0, 0)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while retreiving Itineraries": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": itineraries})
}

func RetrieveItineraries(c *gin.Context, db *sql.DB, limit, offset, referrerId, categoryId int) ([]models.Itinerary, error) {

	itineraries := []models.Itinerary{}

	condition := common.BuildCondition(referrerId, categoryId)

	rows, err := db.QueryContext(c, `SELECT ID,
			category_id,
			referrer_id,
			img,
			title,
			content,
			Created_At,
			Updated_At,
			COUNT(*) OVER()
		FROM Itineraries WHERE `+
		condition+` ORDER BY id ASC `+
		database.FormatLimitOffset(limit, offset))

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		itinerary, err := scanItinerary(rows)

		if err != nil {
			return nil, err
		}

		if itinerary.ID != 0 {

			itineraries = append(itineraries, itinerary)
		}
	}
	defer rows.Close()

	return itineraries, nil
}

func retrieveItinerary(c *gin.Context, referrerId int, categoryId int, db *sql.DB) ([]models.Itinerary, error) {

	itineraries := []models.Itinerary{}

	rows, err := db.QueryContext(c, `SELECT id,
			category_id,
			referrer_id,
			img,
			title,
			content,
			Created_At,
			Updated_At
		
		FROM Itineraries WHERE referrer_id = ? AND category_id`, referrerId, categoryId,
	)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		itinerary, err := scanItinerary(rows)
		itineraries = append(itineraries, itinerary)
		if err != nil {
			return nil, err
		}
	}

	return itineraries, nil
}

func scanItinerary(rows *sql.Rows) (models.Itinerary, error) {
	itinerary := models.Itinerary{}
	var n int

	err := rows.Scan(&itinerary.ID,
		&itinerary.CategoryID,
		&itinerary.ReferrerID,
		&itinerary.Img,
		&itinerary.Title,
		&itinerary.Content,
		(*time.Time)(&itinerary.CreatedAt),
		(*time.Time)(&itinerary.UpdatedAt),
		&n,
	)

	return itinerary, err
}
