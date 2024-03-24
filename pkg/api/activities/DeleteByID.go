package activities

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/api/storage"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/common"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Delete a activity by ID
// @Description Delete a activity by ID
// @Tags activities
// @Accept json
// @Produce json
// @Param id path int true "Activity ID"
// @Success 200 {string} string "Activity deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /activities/{id} [delete]

func DeleteByID(c *gin.Context) {
	activityID, err := strconv.Atoi(c.Param("id"))
	categoryId := common.Activities
	db := database.Database.DB

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	// Delete images from GCS
	err = storage.Delete("ACTIVITY", activityID, categoryId)
	if err != nil {
		if err.Error() == "storage: object doesn't exist" {
			c.JSON(http.StatusOK, gin.H{"message": "Activity deleted successfully"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting images": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM Faqs WHERE Referrer_id = ? AND category_id = ?`, activityID, categoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting Faq": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM Images WHERE Referrer_id = ? AND category_id = ?`, activityID, categoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting SlideImage": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM Itineraries WHERE Referrer_id = ? AND category_id = ?`, activityID, categoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting Itinerary": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM Activities WHERE ID = ?`, activityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting Activity": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity deleted successfully"})
}
