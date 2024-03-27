package contentTypes

import (
	"net/http"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create a new activityType
// @Description Create a new activityType
// @Tags activityTypes
// @Accept json
// @Produce json
// @Param activityType body activityType true "ContentType object"
// @Success 200 {object} activityType "ContentType created successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /activityTypes [post]
func Create(c *gin.Context) {
	var newContentType models.ContentType
	db := database.Database.DB

	result, err := db.ExecContext(c, `
			INSERT INTO ContentTypes (
				title,
				description,
				Created_At,
				Updated_At
			)
        VALUES (?, ?, ?, ?)`,

		newContentType.Title,
		newContentType.Description,
		time.Now(),
		time.Now())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	newContentType.ID = int(id)

	c.JSON(http.StatusOK, gin.H{"data": newContentType})
}
