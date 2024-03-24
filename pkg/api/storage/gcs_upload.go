package storage

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/common"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/models"
)

// Constants for configuration keys
const (
	googleBucketNameKey = "GOOGLE-BUCKET-NAME"
	googleProjectIDKey  = "GOOGLE-PROJECT-ID"
)

// UploadImagesHandler handles image uploads.
func UploadImagesHandler(c *gin.Context) {
	// Extract necessary information from the request
	db := database.Database.DB
	imageType := strings.ToUpper(c.Query("image"))
	productType := strings.ToUpper(c.Query("category"))
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing ID"})
		return
	}

	// Parse multipart form for file uploads
	if err := c.Request.ParseMultipartForm(500000); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get category ID from the common package
	categoryID, err := common.GetCategoryId(productType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get client URI from the request headers
	// clientURI := c.Request.Header.Get("Client-Uri")
	// fmt.Println("Client-Uri: ", clientURI)

	// Retrieve configuration information from the database
	config := database.Database.Config
	bucketName := config[googleBucketNameKey]
	folderName := productType + "/"
	projectID := config[googleProjectIDKey]

	// Variables to store uploaded images and context
	var newGalleryImages []models.GalleryImage
	ctx := context.Background()

	// Create a new Cloud Storage client
	client, err := storage.NewClient(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Cloud Storage client"})
		return
	}
	defer client.Close()

	// Create a handle for the Cloud Storage bucket
	bucket := client.Bucket(bucketName)

	// Check if the bucket exists, create it if not
	if _, err := bucket.Attrs(ctx); err != nil {
		if err == storage.ErrBucketNotExist {
			if err := bucket.Create(ctx, projectID, nil); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bucket"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get bucket"})
			return
		}
	}

	// Retrieve uploaded files from the request
	files := c.Request.MultipartForm.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No files were uploaded"})
		return
	}

	// Delete old references based on the image type
	// deleteOldReferences(c, db, imageType, id, categoryID)

	// Delete old images from Cloud Storage
	// deleteImagesFromGCS(c, productType, id, categoryID)

	// Process each uploaded file
	for _, file := range files {
		strID := strconv.Itoa(id)
		filePath := folderName + strID + "_" + file.Filename
		obj := bucket.Object(filePath)

		// Upload the file to Cloud Storage
		if err := uploadFileToStorage(ctx, obj, file); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Construct the file location URL
		fileLocation := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, filePath)

		// Create a new GalleryImage model
		newGalleryImage := models.GalleryImage{
			CategoryID: strconv.Itoa(categoryID),
			ReferrerID: strID,
			Img:        fileLocation,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		// Append the new image to the slice
		newGalleryImages = append(newGalleryImages, newGalleryImage)

		// Insert the image into the database based on the image type
		err := insertImageIntoDatabase(c, db, imageType, categoryID, id, newGalleryImage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": newGalleryImages})

	}

}

// deleteOldReferences deletes old references of images from the database based on the image type.
func deleteOldReferences(c *gin.Context, db *sql.DB, imageType string, id, categoryID int) {
	var tableName, columnName string

	switch imageType {
	case "GALLERY":
		tableName = "Images"
		columnName = "referrer_id"
	case "SLIDE":
		tableName = "Images"
		columnName = "referrer_id"
	}

	// Construct the SQL query to delete old references
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ? AND category_id = ?", tableName, columnName)
	if _, err := db.ExecContext(c, query, id, categoryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error while deleting %s: %s", imageType, err.Error())})
		return
	}
}

// uploadFileToStorage uploads a file to Cloud Storage.
func uploadFileToStorage(ctx context.Context, obj *storage.ObjectHandle, file *multipart.FileHeader) error {
	w := obj.NewWriter(ctx)
	fileReader, err := file.Open()
	defer fileReader.Close()

	if err != nil {
		return fmt.Errorf("Failed to open file: %w", err)
	}

	if _, err := io.Copy(w, fileReader); err != nil {
		return fmt.Errorf("Failed to upload file: %w", err)
	}

	if err := w.Close(); err != nil {
		return fmt.Errorf("Failed to close writer: %w", err)
	}

	return nil
}

// insertImageIntoDatabase inserts an image into the database based on the image type.
func insertImageIntoDatabase(c *gin.Context, db *sql.DB, imageType string, categoryID, id int, newGalleryImage models.GalleryImage) error {
	var tableName string

	switch imageType {
	case "GALLERY":
		tableName = "Images"
	case "SLIDE":
		tableName = "Images"
	default:
		tableName = "Images"
	}

	// Construct the SQL query to insert the image into the database
	query := fmt.Sprintf("INSERT INTO %s (category_id, referrer_id, img) VALUES (?, ?, ?)", tableName)
	if _, err := db.Exec(query, newGalleryImage.CategoryID, id, newGalleryImage.Img); err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

// deleteImagesFromGCS deletes old images from Cloud Storage.
func deleteImagesFromGCS(c *gin.Context, productType string, id, categoryID int) {
	err := Delete(productType, id, categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while deleting images": err.Error()})
		return
	}
}
