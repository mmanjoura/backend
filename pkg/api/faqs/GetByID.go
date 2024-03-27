package faqs

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Get a faq by ID
// @Description Get a faq by ID
// @Tags faqs
// @Accept json
// @Produce json
// @Param id path int true "Faq ID"
// @Success 200 {object} models.Faq "Successfully retrieved a faq"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Faq not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /faqs/{id} [get]
func GetByID(c *gin.Context) {
	faqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid faq ID"})
		return
	}

	db := database.Database.DB

	var faq models.Faq
	err = db.QueryRowContext(c, `
	SELECT 	ID,
	category_id,
	referrer_id,
	title,
	content,
	collapseTarget,
	Created_At,
	Updated_At
	FROM faqs WHERE ID = ?`, faqID).
		Scan(&faq.ID,
			&faq.CategoryID,
			&faq.ReferrerID,
			&faq.Title,
			&faq.Content,
			&faq.CollapseTarget,
			&faq.CreatedAt,
			&faq.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error while getting faq": "faq not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": faq})
}
