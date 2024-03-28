package comments

import (
	"net/http"
	"strconv"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// UpdateByID godoc
// @Summary Update a comment by ID
// @Description Update a comment by ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Param comment body comment true "Comment object"
// @Success 200 {string} string "Tour updated successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /comments/{id} [put]
func UpdateByID(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	var updatedComment models.Comment

	if err := c.ShouldBindJSON(&updatedComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.Database.DB
	defer db.Close()

	_, err = db.ExecContext(c, `
								UPDATE Comments
								SET subject = ?,
									email = ?,
									message = ?,
									Updated_At = ?
							WHERE id = ?`,
		updatedComment.Subject, updatedComment.Email,
		updatedComment.Message,
		time.Now(), commentID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while updating comment": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tour updated successfully"})
}
