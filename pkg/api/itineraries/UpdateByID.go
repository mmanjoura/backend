package itineraries

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/common"
	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// UpdateByID godoc
// @Summary Update a itinerary by ID
// @Description Update a itinerary by ID
// @Tags itineraries
// @Accept json
// @Produce json
// @Param id path int true "Tour ID"
// @Param itinerary body models.Tour true "Tour object"
// @Success 200 {string} string "Tour updated successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /itineraries/{id} [put]
func UpdateByID(c *gin.Context) {

	var updatedItinerary models.Itinerary
	db := database.Database.DB

	id := strings.ToUpper(c.Query("id"))
	productType := strings.ToUpper(c.Query("category"))

	categoryId, err := common.GetCategoryId(productType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// _, err = db.ExecContext(c, `DELETE FROM Itineraries WHERE referrer_id = ? AND category_id = ?`, id, categoryId)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error while deleting Itinerary": err.Error()})
	// 	return
	// }

	if err := c.ShouldBindJSON(&updatedItinerary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = db.ExecContext(c, `
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
		updatedItinerary.Img,
		updatedItinerary.Title,
		updatedItinerary.Content,
		time.Now(),
		time.Now())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while updating itinerary": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Itineraruy updated successfully"})
}

func insertItinerary(c *gin.Context, db *sql.DB, id, categoryId int, updatedItinerary models.Itinerary) error {
	_, err := db.ExecContext(c, `
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
		updatedItinerary.Img,
		updatedItinerary.Title,
		updatedItinerary.Content,
		time.Now(),
		time.Now())

	return err
}
