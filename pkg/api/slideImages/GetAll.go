package slideImages

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

func GetAll(c *gin.Context) {
	db := database.Database.DB
	limit, offset := common.GetPaginationParams(c)
	faqs, err := RetrieveSlideImages(c, db, limit, offset, 0, 0)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while retreiving slideImages": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": faqs})
}

func RetrieveSlideImages(c *gin.Context, db *sql.DB, limit, offset, referrerId, categoryId int) ([]models.SlideImage, error) {

	condition := common.BuildCondition(referrerId, categoryId)

	slideImages := []models.SlideImage{}

	rows, err := db.QueryContext(c, `SELECT ID,
			category_id,
			referrer_id,
			img,
			Created_At,
			Updated_At,
			COUNT(*) OVER()
			FROM SlideImages WHERE `+
		condition+` ORDER BY id ASC `+
		database.FormatLimitOffset(limit, offset))

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		slideImage, err := scanSlideImages(rows)

		if err != nil {
			return nil, err
		}

		if slideImage.ID != 0 {

			slideImages = append(slideImages, slideImage)
		}
	}
	defer rows.Close()

	return slideImages, nil
}

func scanSlideImages(rows *sql.Rows) (models.SlideImage, error) {
	slideImage := models.SlideImage{}
	var n int

	err := rows.Scan(&slideImage.ID,
		&slideImage.CategoryID,
		&slideImage.ReferrerID,
		&slideImage.Img,
		(*time.Time)(&slideImage.CreatedAt),
		(*time.Time)(&slideImage.UpdatedAt),
		&n,
	)

	return slideImage, err
}
