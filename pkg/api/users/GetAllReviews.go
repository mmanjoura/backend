package users

import (
	"database/sql"
	"net/http"
	"time"

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
// @Router /reviews [get]
func Healthcheck(g *gin.Context) {
	g.JSON(http.StatusOK, "ok")
}

// GetAllReviews godoc
// @Summary   User reviews List
// @Tags      user
// @Accept    json
// @Produce   json
// @Success   200  {array}   success
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  UserAuth
// @Router    /users/reviews [get]
func GetAllReviews(c *gin.Context) {
	offset, limit := common.GetPaginationParams(c)
	reviews, err := retrieveReviews(c, offset, limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while retreiving Reviews": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

func retrieveReviews(c *gin.Context, offset, limit int) ([]models.Review, error) {
	db := database.Database.DB

	reviews := []models.Review{}

	rows, err := db.QueryContext(c, `SELECT id,
				text,
				rating,
				product_ID,
				user_ID,
				username,
				Created_At,
				Updated_At
			FROM Reviews,
			COUNT(*) OVER()
		FROM Reviews WHERE 1 = 1
		ORDER BY id ASC `+database.FormatLimitOffset(limit, offset),
	)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		review, err := scanReview(rows)

		if err != nil {
			return nil, err
		}

		if review.ID != 0 {

			reviews = append(reviews, review)
		}
	}

	return reviews, nil
}

func scanReview(rows *sql.Rows) (models.Review, error) {
	review := models.Review{}
	var n int
	err := rows.Scan(&review.ID,
		&review.UserID,
		&review.ProductID,
		&review.Text,
		&review.Rating,
		&review.Username,
		(*time.Time)(&review.CreatedAt),
		(*time.Time)(&review.UpdatedAt),
		&n,
	)

	return review, err
}
