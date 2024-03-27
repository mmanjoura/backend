package contentTypes

import (
	"net/http"
	"strconv"

	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Get a contentType by ID
// @Description Get a contentType by ID
// @Tags contentTypes
// @Accept json
// @Produce json
// @Param id path int true "ContentType ID"
// @Success 200 {object} contentType "ContentType found"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /contentTypes/{id} [get]
func GetByID(c *gin.Context) {
	contentTypeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contentType ID"})
		return
	}

	db := database.Database.DB

	var contentType models.ContentType
	err = db.QueryRowContext(c, `
	SELECT 	ID,
	Name,
	Description,
	Created_At,
	Updated_At
	FROM contentTypes WHERE ID = ?`, contentTypeID).
		Scan(&contentType.ID,
			&contentType.Title,
			&contentType.Description,
			&contentType.CreatedAt,
			&contentType.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error while getting contentType": "contentType not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contentType})
}
