package galleryImages

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
// @Summary Get all galleryImages
// @Description Get all galleryImages
// @Tags galleryImages
// @Accept json
// @Produce json
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Success 200 {object} models.GalleryImage "Successfully retrieved galleryImages"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /galleryImages [get]

func GetAll(c *gin.Context) {
	db := database.Database.DB
	offset, limit := common.GetPaginationParams(c)
	galleryImages, err := RetrieveImages(c, db, limit, offset, 0, 0)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while retreiving galleryImages": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": galleryImages})
}

func RetrieveImages(c *gin.Context, db *sql.DB, limit, offset, referrerId, categoryId int) ([]models.Image, error) {

	condition := common.BuildCondition(referrerId, categoryId)

	tourImages := []models.Image{}

	rows, err := db.QueryContext(c, `SELECT ID,
			category_id,
			referrer_id,
			img,
			Created_At,
			Updated_At,
			COUNT(*) OVER()
		FROM images WHERE `+
		condition+` ORDER BY id ASC `+
		database.FormatLimitOffset(limit, offset))

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tourImage, err := scanTourImages(rows)

		if err != nil {
			return nil, err
		}

		if tourImage.ID != 0 {

			tourImages = append(tourImages, tourImage)
		}
	}
	defer rows.Close()

	return tourImages, nil
}

func scanTourImages(rows *sql.Rows) (models.Image, error) {
	tourImage := models.Image{}
	var n int

	err := rows.Scan(&tourImage.ID,
		&tourImage.CategoryID,
		&tourImage.ReferrerID,
		&tourImage.Img,
		(*time.Time)(&tourImage.CreatedAt),
		(*time.Time)(&tourImage.UpdatedAt),
		&n,
	)

	return tourImage, err
}
