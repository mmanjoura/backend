package itineraries

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
// @Summary Create a new itinerary
// @Description Create a new itinerary
// @Tags itineraries
// @Accept json
// @Produce json
// @Param itinerary body models.Itinerary true "Itinerary object"
// @Success 200 {object} models.Itinerary "Successfully created a new itinerary"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /itineraries [post]
func Create(c *gin.Context) {
	var newItinerary models.CreateItinerary

	if err := c.ShouldBindJSON(&newItinerary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := strings.ToUpper(c.Query("id"))
	category := strings.ToUpper(c.Query("category"))

	if id == "" || category == "" {
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
		INSERT INTO Itineraries (
			category_id,
			referrer_id,
			img,
			title,
			content,
			Created_At,
			Updated_At)
        VALUES (?, ?, ?, ?, ?, ?, ?)`,

		categoryId,
		id,
		newItinerary.Img,
		newItinerary.Title,
		newItinerary.Content,
		time.Now(),
		time.Now())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
