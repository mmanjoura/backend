package tours

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage/backend/pkg/api/storage"
	"github.com/mmanjoura/niya-voyage/backend/pkg/common"
	"github.com/mmanjoura/niya-voyage/backend/pkg/database"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Delete a tour by ID
// @Description Delete a tour by ID
// @Tags tours
// @Accept json
// @Produce json
// @Param id path int true "Tour ID"
// @Success 200 {string} string "Tour deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tours/{id} [delete]

func DeleteByID(c *gin.Context) {
	tourID, err := strconv.Atoi(c.Param("id"))
	categoryId := common.Tours
	db := database.Database.DB

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tour ID"})
		return
	}

	// Delete images from GCS
	err = storage.Delete("TOUR", tourID, categoryId)
	if err != nil {
		if err.Error() == "storage: object doesn't exist" {
			c.JSON(http.StatusOK, gin.H{"message": "Tour deleted successfully"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting images": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM Faqs WHERE Referrer_id = ? AND category_id = ?`, tourID, categoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting Faq": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM Images WHERE Referrer_id = ? AND category_id = ?`, tourID, categoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting Images": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM SlideImages WHERE Referrer_id = ? AND category_id = ?`, tourID, categoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting SlideImage": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM GalleryImages WHERE Referrer_id = ? AND category_id = ?`, tourID, categoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting GalleryImages": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM Itineraries WHERE Referrer_id = ? AND category_id = ?`, tourID, categoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting Itinerary": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM Tours WHERE ID = ?`, tourID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting Tour": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tour deleted successfully"})
}
