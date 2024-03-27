package golfs

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage/backend/pkg/api/storage"
	"github.com/mmanjoura/niya-voyage/backend/pkg/common"
	"github.com/mmanjoura/niya-voyage/backend/pkg/database"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Delete a golf by ID
// @Description Delete a golf by ID
// @Tags golfs
// @Accept json
// @Produce json
// @Param id path int true "Golf ID"
// @Success 200 {string} string "Golf deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /golfs/{id} [delete]

func DeleteByID(c *gin.Context) {
	golfID, err := strconv.Atoi(c.Param("id"))
	categoryId := common.Activities
	db := database.Database.DB

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid golf ID"})
		return
	}

	// Delete images from GCS
	err = storage.Delete("GOLF", golfID, categoryId)
	if err != nil {
		if err.Error() == "storage: object doesn't exist" {
			c.JSON(http.StatusOK, gin.H{"message": "Golf deleted successfully"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting images": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM Faqs WHERE Referrer_id = ? AND category_id = ?`, golfID, categoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting Faq": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM Images WHERE Referrer_id = ? AND category_id = ?`, golfID, categoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting SlideImage": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM Itineraries WHERE Referrer_id = ? AND category_id = ?`, golfID, categoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting Itinerary": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM Golfs WHERE ID = ?`, golfID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting Golf": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Golf deleted successfully"})
}
