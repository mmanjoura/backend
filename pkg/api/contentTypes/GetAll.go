package contentTypes

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
// @Summary 		Get all ContentTypes
// @Description 	Get all ContentTypes
// @Tags 			ContentTypes
// @Accept  		json
// @Produce 		json
// @Param 			offset query int false "Offset"
// @Param 			limit query int false "Limit"
// @Success 		200 {object} []ContentType
// @Failure 		400 {string} string "Bad Request"
// @Failure 		500 {string} string "Internal Server Error"
// @Router 			/contentTypes [get]
func GetAll(c *gin.Context) {
	db := database.Database.DB
	offset, limit := common.GetPaginationParams(c)
	contentTypes, err := RetrieveContentTypes(c, db, limit, offset)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while retreiving ContentTypes": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contentTypes})
}

func RetrieveContentTypes(c *gin.Context, db *sql.DB, limit, offset int) ([]models.ContentType, error) {

	// Build WHERE clause. Each part of the clause is AND-ed together to further
	// restrict the results. Placeholders are added to "args" and are used
	// to avoid SQL injection.
	// Each filter field is optional.
	_, args := []string{"1 = 1"}, []interface{}{}

	contentTypes := []models.ContentType{}

	rows, err := db.QueryContext(c, `SELECT ID,
			title,
			description,
			Created_At,
			Updated_At,
			COUNT(*) OVER()
		FROM ContentTypes `,
		args...,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		contentType, err := scanContentType(rows)

		if err != nil {
			return nil, err
		}

		if contentType.ID != 0 {

			contentTypes = append(contentTypes, contentType)
		}
	}
	defer rows.Close()

	return contentTypes, nil
}

func scanContentType(rows *sql.Rows) (models.ContentType, error) {
	contentType := models.ContentType{}
	var n int

	err := rows.Scan(&contentType.ID,
		&contentType.Title,
		&contentType.Description,
		(*time.Time)(&contentType.CreatedAt),
		(*time.Time)(&contentType.UpdatedAt),
		&n,
	)

	return contentType, err
}
