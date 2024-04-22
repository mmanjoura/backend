package storage

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
)

// Constants for configuration keys
const (
	backupBucket = "BACKUPS-GOOGLE-BUCKET"
	googleProjectIDKey  = "GOOGLE-PROJECT-ID"
)

// UploadImagesHandler handles image uploads.
func BackUp(c *gin.Context) {

	config := database.Database.Config
	bucketName := config[backupBucket]
	projectID := config[googleProjectIDKey]
	folderName := "DB-BACKUPS/"

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

	// Create a new object in the bucket
	obj := bucket.Object(folderName + time.Now().Format("20060102150405") + "_backup.db")
	w := obj.NewWriter(ctx)

	// Create an io.reader from the db file
	file, err := os.Open(config["DB_NAME"])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to open file: %w", err).Error()})
		return
	}
	defer file.Close()

	if _, err := io.Copy(w, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to upload file: %w", err).Error()})
		return
	}

	if err := w.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to close writer: %w", err).Error()})
		return
	}


}
