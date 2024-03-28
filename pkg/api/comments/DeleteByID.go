package comments

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Delete a comment by ID
// @Description Delete a comment by ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {string} string "Comment deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /comments/{id} [delete]
func DeleteByID(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("id"))
	db := database.Database.DB
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM comments WHERE ID = ?`, commentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting comment": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "comment deleted successfully"})
}
