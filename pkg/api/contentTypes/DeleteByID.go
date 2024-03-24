package contentTypes

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Delete a activityType by ID
// @Description Delete a activityType by ID
// @Tags activityTypes
// @Accept json
// @Produce json
// @Param id path int true "ContentType ID"
// @Success 200 {string} string "ContentType deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /activityTypes/{id} [delete]
func DeleteByID(c *gin.Context) {
	activityTypeId, err := strconv.Atoi(c.Param("id"))
	db := database.Database.DB
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tour ID"})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM ContentTypes WHERE ID = ?`, activityTypeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting ContentType": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ContentTypes deleted successfully"})
}
