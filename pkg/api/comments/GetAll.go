package comments

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
// @Summary 		Get all Comments
// @Description 	Get all Comments
// @Tags 			Comments
// @Accept  		json
// @Produce 		json
// @Param 			offset query int false "Offset"
// @Param 			limit query int false "Limit"
// @Success 		200 {object} []Comment
// @Failure 		400 {string} string "Bad Request"
// @Failure 		500 {string} string "Internal Server Error"
// @Router 			/comments [get]
func GetAll(c *gin.Context) {
	db := database.Database.DB
	offset, limit := common.GetPaginationParams(c)
	comments, err := RetrieveComments(c, db, limit, offset)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while retreiving Comments": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

func RetrieveComments(c *gin.Context, db *sql.DB, limit, offset int) ([]models.Comment, error) {

	// Build WHERE clause. Each part of the clause is AND-ed together to further
	// restrict the results. Placeholders are added to "args" and are used
	// to avoid SQL injection.
	// Each filter field is optional.
	_, args := []string{"1 = 1"}, []interface{}{}

	comments := []models.Comment{}

	rows, err := db.QueryContext(c, `SELECT ID,
			subject,
			email,
			message,
			Created_At,
			Updated_At,
			COUNT(*) OVER()
		FROM Comments `,
		args...,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		comment, err := scanComment(rows)

		if err != nil {
			return nil, err
		}

		if comment.ID != 0 {

			comments = append(comments, comment)
		}
	}
	defer rows.Close()

	return comments, nil
}

func scanComment(rows *sql.Rows) (models.Comment, error) {
	comment := models.Comment{}
	var n int

	err := rows.Scan(&comment.ID,
		&comment.Subject,
		&comment.Email,
		&comment.Message,
		(*time.Time)(&comment.CreatedAt),
		(*time.Time)(&comment.UpdatedAt),
		&n,
	)

	return comment, err
}
