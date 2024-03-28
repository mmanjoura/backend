package comments

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Get a comment by ID
// @Description Get a comment by ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {object} comment "Comment found"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /comments/{id} [get]
func GetByID(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	db := database.Database.DB

	var comment models.Comment
	err = db.QueryRowContext(c, `
	SELECT 	ID,
	subject,
	email,
	message,
	Created_At,
	Updated_At
	FROM comments WHERE ID = ?`, commentID).
		Scan(&comment.ID,
			&comment.Subject,
			&comment.Email,
			&comment.Message,
			&comment.CreatedAt,
			&comment.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error while getting comment": "comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}
