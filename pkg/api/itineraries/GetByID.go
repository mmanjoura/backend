package itineraries

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Get a itinerary by ID
// @Description Get a itinerary by ID
// @Tags itinerarys
// @Accept json
// @Produce json
// @Param id path int true "Itinerary ID"
// @Success 200 {object} models.Itinerary "Successfully retrieved a itinerary"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Itinerary not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /itinerarys/{id} [get]
func GetByID(c *gin.Context) {
	referrerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid itinerary ID"})
		return
	}

	db := database.Database.DB

	var itinerary models.Itinerary
	err = db.QueryRowContext(c, `
	SELECT 	ID,
    category_id,
	referrer_id,
	img,
	title,
	content,
	Created_At,
	Updated_At
	FROM itineraries WHERE referrer_id = ?`, referrerId).
		Scan(&itinerary.ID,
			&itinerary.CategoryID,
			&itinerary.ReferrerID,
			&itinerary.Img,
			&itinerary.Title,
			&itinerary.Content,
			&itinerary.CreatedAt,
			&itinerary.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error while getting Itinerary": "Itinerary not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": itinerary})
}
