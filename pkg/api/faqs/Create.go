package faqs

import (
	"net/http"
	"strings"
	"time"

	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/common"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create a new FAQ
// @Description Create a new FAQ
// @Tags faqs
// @Accept json
// @Produce json
// @Param faq body models.Faq true "FAQ object"
// @Success 200 {object} models.Faq "Successfully created a FAQ"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /faqs [post]
func Create(c *gin.Context) {
	var newFaq models.Faq

	if err := c.ShouldBindJSON(&newFaq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activityId := strings.ToUpper(c.Query("id"))
	category := strings.ToUpper(c.Query("category"))

	if activityId == "" || category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id and category are required"})
		return

	}

	categoryId, err := common.GetCategoryId(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.Database.DB

	result, err := db.ExecContext(c, `
			INSERT INTO Faqs (
				category_id,
				referrer_id,
				title,
				content,
				collapseTarget,
				Created_At,
				Updated_At
			)
        VALUES (?, ?, ?, ?, ?, ?, ?)`,

		categoryId,
		activityId,
		newFaq.Title,
		newFaq.Content,
		newFaq.CollapseTarget,
		time.Now(),
		time.Now())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	newFaq.ID = int(id)

	c.JSON(http.StatusOK, gin.H{"data": newFaq})
}
