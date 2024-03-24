package images

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/common"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/models"

	"github.com/gin-gonic/gin"
)

// GetAll 			godoc

func GetAll(c *gin.Context) {
	db := database.Database.DB
	limit, offset := common.GetPaginationParams(c)
	images, err := RetrieveImages(c, db, limit, offset, 0, 0)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while retreiving images": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": images})
}

func RetrieveImages(c *gin.Context, db *sql.DB, limit, offset, referrerId, categoryId int) ([]models.Image, error) {

	condition := common.BuildCondition(referrerId, categoryId)

	Images := []models.Image{}

	rows, err := db.QueryContext(c, `SELECT ID,
			category_id,
			referrer_id,
			img,
			Created_At,
			Updated_At,
			COUNT(*) OVER()
			FROM Images WHERE `+
		condition+` ORDER BY id ASC `+
		database.FormatLimitOffset(limit, offset))

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		image, err := scanImages(rows)

		if err != nil {
			return nil, err
		}

		if image.ID != 0 {

			Images = append(Images, image)
		}
	}
	defer rows.Close()

	return Images, nil
}

func scanImages(rows *sql.Rows) (models.Image, error) {
	image := models.Image{}
	var n int

	err := rows.Scan(&image.ID,
		&image.CategoryID,
		&image.ReferrerID,
		&image.Img,
		(*time.Time)(&image.CreatedAt),
		(*time.Time)(&image.UpdatedAt),
		&n,
	)

	return image, err
}
