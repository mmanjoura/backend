package images

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

func RetrieveGalleryImages(c *gin.Context, db *sql.DB, limit, offset, referrerId, categoryId int) ([]models.GalleryImage, error) {

	condition := common.BuildCondition(referrerId, categoryId)

	galleryImages := []models.GalleryImage{}

	rows, err := db.QueryContext(c, `SELECT ID,
			category_id,
			referrer_id,
			img,
			Created_At,
			Updated_At,
			COUNT(*) OVER()
			FROM GalleryImages WHERE `+
		condition+` ORDER BY id ASC `+
		database.FormatLimitOffset(limit, offset))

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		galleryImage, err := scanGalleryImages(rows)

		if err != nil {
			return nil, err
		}

		if galleryImage.ID != 0 {

			galleryImages = append(galleryImages, galleryImage)
		}
	}
	defer rows.Close()

	return galleryImages, nil
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

func scanGalleryImages(rows *sql.Rows) (models.GalleryImage, error) {
	galleryImage := models.GalleryImage{}
	var n int

	err := rows.Scan(&galleryImage.ID,
		&galleryImage.CategoryID,
		&galleryImage.ReferrerID,
		&galleryImage.Img,
		(*time.Time)(&galleryImage.CreatedAt),
		(*time.Time)(&galleryImage.UpdatedAt),
		&n,
	)

	return galleryImage, err
}
