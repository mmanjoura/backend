package comments

import (
	"net/http"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create a new FAQ
// @Description Create a new FAQ
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body models.Comment true "FAQ object"
// @Success 200 {object} models.Comment "Successfully created a FAQ"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /comments [post]
func Create(c *gin.Context) {
	var newComment models.Comment

	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.Database.DB

	result, err := db.ExecContext(c, `
			INSERT INTO comments (
				subject,
				email,
				message,
				Created_At,
				Updated_At
			)
        VALUES ( ?, ?, ?, ?, ?)`,

		newComment.Subject,
		newComment.Email,
		newComment.Message,
		time.Now(),
		time.Now())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	newComment.ID = int(id)

	c.JSON(http.StatusOK, gin.H{"data": newComment})
}
