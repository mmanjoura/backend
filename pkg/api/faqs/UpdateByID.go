package faqs

import (
	"net/http"
	"strings"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/common"
	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// UpdateByID godoc
// @Summary Update a faq by ID
// @Description Update a faq by ID
// @Tags faqs
// @Accept json
// @Produce json
// @Param id path int true "Tour ID"
// @Param faq body models.Tour true "Tour object"
// @Success 200 {string} string "Tour updated successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /faqs/{id} [put]
func UpdateByID(c *gin.Context) {
	var updatedFaq models.Faq
	db := database.Database.DB

	id := strings.ToUpper(c.Query("id"))
	productType := strings.ToUpper(c.Query("category"))

	categoryId, err := common.GetCategoryId(productType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `DELETE FROM faqs WHERE referrer_id = ? AND category_id = ?`, id, categoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting Itinerary": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&updatedFaq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `
		INSERT INTO Faqs (
			category_id,
			referrer_id,
			title,
			content,
			collapseTarget,
			Created_At,
			Updated_At)
        VALUES (?, ?, ?, ?, ?, ?, ?)`,

		categoryId,
		id,
		updatedFaq.Title,
		updatedFaq.Content,
		updatedFaq.CollapseTarget,
		time.Now(),
		time.Now())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while updating itinerary": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tour updated successfully"})
}
