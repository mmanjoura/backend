package faqs

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
// @Summary			Get all faqs with pagination
// @Description 	Get a list of all faqs with optional pagination
// @Tags 			faqs
// @Security 		ApiKeyAuth
// @Produce 		json
// @Param 			offset query int false "Offset for pagination" default(0)
// @Param 			limit query int false "Limit for pagination" default(10)
// @Success 		200 {array} models.Faq "Successfully retrieved list of faqs"
// @Router 			/faqs [get]
func GetAll(c *gin.Context) {
	db := database.Database.DB
	limit, offset := common.GetPaginationParams(c)
	faqs, err := RetrieveFaqs(c, db, limit, offset, 0, 0)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while retreiving Faqs": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": faqs})
}

func RetrieveFaqs(c *gin.Context, db *sql.DB, limit, offset int, referrerId, categoryId int) ([]models.Faq, error) {

	condition := common.BuildCondition(referrerId, categoryId)

	faqs := []models.Faq{}

	rows, err := db.QueryContext(c, `SELECT ID,
			category_id,
			referrer_id,
			title,
			content,
			collapseTarget,
			Created_At,
			Updated_At,
			COUNT(*) OVER()
		FROM Faqs WHERE  `+
		condition+` ORDER BY id ASC `+
		database.FormatLimitOffset(limit, offset))

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		faq, err := scanFaq(rows)

		if err != nil {
			return nil, err
		}

		if faq.ID != 0 {

			faqs = append(faqs, faq)
		}
	}
	defer rows.Close()

	return faqs, nil
}

func scanFaq(rows *sql.Rows) (models.Faq, error) {
	faq := models.Faq{}
	var n int

	err := rows.Scan(&faq.ID,
		&faq.CategoryID,
		&faq.ReferrerID,
		&faq.Title,
		&faq.Content,
		&faq.CollapseTarget,
		(*time.Time)(&faq.CreatedAt),
		(*time.Time)(&faq.UpdatedAt),
		&n,
	)

	return faq, err
}
