package contentTypes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/models"

	"github.com/gin-gonic/gin"
)

// UpdateByID godoc
// @Summary Update a contentType by ID
// @Description Update a contentType by ID
// @Tags contentTypes
// @Accept json
// @Produce json
// @Param id path int true "ContentType ID"
// @Param contentType body contentType true "ContentType object"
// @Success 200 {string} string "Tour updated successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /contentTypes/{id} [put]
func UpdateByID(c *gin.Context) {
	contentTypeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contentType ID"})
		return
	}

	var updatedContentType models.ContentType

	if err := c.ShouldBindJSON(&updatedContentType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.Database.DB
	defer db.Close()

	_, err = db.ExecContext(c, `
								UPDATE ContentTypes
								SET name = ?,
									description = ?,
									Updated_At = ?
							WHERE id = ?`,
		updatedContentType.Title,
		updatedContentType.Description,
		time.Now(), contentTypeID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while updating contentType": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tour updated successfully"})
}
