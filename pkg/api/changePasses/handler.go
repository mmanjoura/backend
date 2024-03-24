package changePasses

// import (
// 	"encoding/json"
// 	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/cache"
// 	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"
// 	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/models"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// // @BasePath /api/v1

// // Healthcheck godoc
// // @Summary ping example
// // @Schemes
// // @Description do ping
// // @Tags example
// // @Accept json
// // @Produce json
// // @Success 200 {string} ok
// // @Router / [get]
// func Healthcheck(g *gin.Context) {
// 	g.JSON(http.StatusOK, "ok")
// }

// // FindPasswords godoc
// // @Summary Get all Passwords with pagination
// // @Description Get a list of all Passwords with optional pagination
// // @Tags Passwords
// // @Security ApiKeyAuth
// // @Produce json
// // @Param offset query int false "Offset for pagination" default(0)
// // @Param limit query int false "Limit for pagination" default(10)
// // @Success 200 {array} models.Password "Successfully retrieved list of Passwords"
// // @Router /Passwords [get]
// func FindPasswords(c *gin.Context) {
// 	var Passwords []models.ChangePass

// 	// Get query params
// 	offsetQuery := c.DefaultQuery("offset", "0")
// 	limitQuery := c.DefaultQuery("limit", "10")

// 	// Convert query params to integers
// 	offset, err := strconv.Atoi(offsetQuery)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset format"})
// 		return
// 	}

// 	limit, err := strconv.Atoi(limitQuery)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit format"})
// 		return
// 	}

// 	// Create a cache key based on query params
// 	cacheKey := "Passwords_offset_" + offsetQuery + "_limit_" + limitQuery

// 	// Try fetching the data from Redis first
// 	cachedPasswords, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
// 	if err == nil {
// 		err := json.Unmarshal([]byte(cachedPasswords), &Passwords)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"data": Passwords})
// 		return
// 	}

// 	// If cache missed, fetch data from the database
// 	database.Database.DB.Offset(offset).Limit(limit).Find(&Passwords)

// 	// Serialize Passwords object and store it in Redis
// 	// serializedPasswords, err := json.Marshal(Passwords)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
// 	// 	return
// 	// }
// 	// err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedPasswords, time.Minute).Err() // Here TTL is set to one hour
// 	// if err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
// 	// 	return
// 	// }

// 	c.JSON(http.StatusOK, gin.H{"data": Passwords})
// }

// // CreatePassword godoc
// // @Summary Create a new Password
// // @Description Create a new Password with the given input data
// // @Tags Passwords
// // @Security ApiKeyAuth
// // @Security JwtAuth
// // @Accept  json
// // @Produce  json
// // @Param   input     body   models.CreatePassword   true   "Create Password object"
// // @Success 201 {object} models.Password "Successfully created Password"
// // @Failure 400 {string} string "Bad Request"
// // @Failure 401 {string} string "Unauthorized"
// // @Router /Passwords [post]
// func CreatePassword(c *gin.Context) {
// 	var input models.CreatePassword

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	Password := models.ChangePass{
// 		MerchantID:      input.MerchantID,
// 		CurrentPassword: input.CurrentPassword,
// 		NewPassword:     input.NewPassword,
// 	}

// 	database.Database.DB.Create(&Password)

// 	// Invalidate cache
// 	keysPattern := "Passwords_offset_*"
// 	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
// 	if err == nil {
// 		for _, key := range keys {
// 			cache.Rdb.Del(cache.Ctx, key)
// 		}
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"data": Password})
// }

// // FindPassword godoc
// // @Summary Find a Password by ID
// // @Description Get details of a Password by its ID
// // @Tags Passwords
// // @Security ApiKeyAuth
// // @Produce json
// // @Param id path string true "Password ID"
// // @Success 200 {object} models.Password "Successfully retrieved Password"
// // @Failure 404 {string} string "Password not found"
// // @Router /Passwords/{id} [get]
// func FindPassword(c *gin.Context) {
// 	var Password models.ChangePass

// 	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Password).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Password not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": Password})
// }

// // UpdatePassword godoc
// // @Summary Update a Password by ID
// // @Description Update the Password details for the given ID
// // @Tags Passwords
// // @Security ApiKeyAuth
// // @Accept  json
// // @Produce  json
// // @Param id path string true "Password ID"
// // @Param input body models.UpdatePassword true "Update Password object"
// // @Success 200 {object} models.Password "Successfully updated Password"
// // @Failure 400 {string} string "Bad Request"
// // @Failure 404 {string} string "Password not found"
// // @Router /Passwords/{id} [put]
// func UpdatePassword(c *gin.Context) {
// 	var Password models.ChangePass
// 	var input models.UpdatePassword

// 	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Password).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Password not found"})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	database.Database.DB.Model(&Password).Updates(models.ChangePass{
// 		MerchantID:      input.MerchantID,
// 		CurrentPassword: input.CurrentPassword,
// 		NewPassword:     input.NewPassword,
// 	})

// 	c.JSON(http.StatusOK, gin.H{"data": Password})
// }

// // DeletePassword godoc
// // @Summary Delete a Password by ID
// // @Description Delete the Password with the given ID
// // @Tags Passwords
// // @Security ApiKeyAuth
// // @Produce json
// // @Param id path string true "Password ID"
// // @Success 204 {string} string "Successfully deleted Password"
// // @Failure 404 {string} string "Password not found"
// // @Router /Passwords/{id} [delete]
// func DeletePassword(c *gin.Context) {
// 	var Password models.ChangePass

// 	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Password).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Password not found"})
// 		return
// 	}

// 	database.Database.DB.Delete(&Password)

// 	c.JSON(http.StatusNoContent, gin.H{"data": true})
// }
