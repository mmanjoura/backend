package storage

import (
	"context"
	"database/sql"
	"fmt"
	"image"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/mmanjoura/niya-voyage/backend/pkg/common"
	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"
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
	var newSlideImages []models.SlideImage
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

	// Process each uploaded file
	for _, file := range files {
		strID := strconv.Itoa(id)

		// convert the multipart file into io.Reader
		fileReader, err := common.FileHeaderToReader(file)

		// decode the multipart file into an image
		img, _, err := image.Decode(fileReader)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		slideImage, err := common.ProcessImage(img, 300, 300)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		galleryImage, err := common.ProcessImage(img, 600, 500)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// create a map of img.Image and string
		imgMap := make(map[string]image.Image)
		imgMap["slideImage"] = slideImage
		imgMap["galleryImage"] = galleryImage

		for key, scaledImage := range imgMap {
			filePath := folderName + key + "_" + strID + "_" + file.Filename
			obj := bucket.Object(filePath)
			imgReader, err := common.ImageToReader(scaledImage)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			// Upload the galleryImag to Cloud Storage
			if err := uploadFileToStorage(ctx, obj, imgReader); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			// Construct the file location URL
			fileLocation := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, filePath)
			// Create a new GalleryImage model

			if strings.ToUpper(key) == "GALLERYIMAGE" {
				newGalleryImage := models.GalleryImage{
					CategoryID: strconv.Itoa(categoryID),
					ReferrerID: strID,
					Img:        fileLocation,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}
				// Append the new image to the slice
				newGalleryImages = append(newGalleryImages, newGalleryImage)
				err = insertImageIntoDatabase(c, db, "GALLERY", categoryID, id, newGalleryImage)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

			}
			if strings.ToUpper(key) == "SLIDEIMAGE" {
				newSlideImage := models.SlideImage{
					CategoryID: strconv.Itoa(categoryID),
					ReferrerID: strID,
					Img:        fileLocation,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}
				// Append the new image to the slice
				newSlideImages = append(newSlideImages, newSlideImage)
				err = insertImageIntoDatabase(c, db, "SLIDE", categoryID, id, newSlideImage)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

			}

		}
	}

}

// deleteOldReferences deletes old references of images from the database based on the image type.
func deleteOldReferences(c *gin.Context, db *sql.DB, imageType string, id, categoryID int) {
	var tableName, columnName string

	switch imageType {
	case "GALLERY":
		tableName = "GalleryImages"
		columnName = "referrer_id"
	case "SLIDE":
		tableName = "SlideImages"
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
func uploadFileToStorage(ctx context.Context, obj *storage.ObjectHandle, img io.Reader) error {
	w := obj.NewWriter(ctx)

	// Convert img into io.reader

	if _, err := io.Copy(w, img); err != nil {
		return fmt.Errorf("Failed to upload file: %w", err)
	}

	if err := w.Close(); err != nil {
		return fmt.Errorf("Failed to close writer: %w", err)
	}

	return nil
}

// insertImageIntoDatabase inserts an image into the database based on the image type.
func insertImageIntoDatabase(c *gin.Context, db *sql.DB, imageType string, categoryID, id int, images interface{}) error {
	var tableName string

	switch imageType {
	case "GALLERY":
		tableName = "GalleryImages"
	case "SLIDE":
		tableName = "SlideImages"
	default:
		tableName = "Images"
	}

	galleryImage, ok := images.(models.GalleryImage)

	if ok {
		// Construct the SQL query to insert the image into the database
		query := fmt.Sprintf("INSERT INTO %s (category_id, referrer_id, img) VALUES (?, ?, ?)", tableName)
		if _, err := db.Exec(query, galleryImage.CategoryID, id, galleryImage.Img); err != nil {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return err
		}
	}
	slideImage, ok := images.(models.SlideImage)
	if ok {
		// Construct the SQL query to insert the image into the database
		query := fmt.Sprintf("INSERT INTO %s (category_id, referrer_id, img) VALUES (?, ?, ?)", tableName)
		if _, err := db.Exec(query, slideImage.CategoryID, id, slideImage.Img); err != nil {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return err
		}

		return nil
	}
	return nil
}
