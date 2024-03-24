package storage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/api/images"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"
)

// Constants for configuration keys
const (
	googleStorage = "GOOGLE-STORAGE"
)

// UploadImages handles image uploads
func Delete(productType string, referrerId, categoryId int) error {

	config := database.Database.Config
	db := database.Database.DB

	bucketName := config[googleBucketNameKey]
	googleStorage := config[googleStorage]

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	ginContext := &gin.Context{}

	tourImages, err := images.RetrieveImages(ginContext, db, 10, 0, referrerId, categoryId)

	for _, tourImage := range tourImages {
		// Create a Cloud Storage object handle

		imageName, err := extractFolderAndImageName(googleStorage, bucketName, tourImage.Img)
		if err != nil {
			return err
		}

		obj := client.Bucket(bucketName).Object(imageName)

		// Delete the object
		if err := obj.Delete(ctx); err != nil {
			return err
		}
	}

	return nil

}

func deleteBucket(googleStorage, bucketName string) error {
	// bucketName := "bucket-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient:", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	bucket := client.Bucket(bucketName)

	if err := bucket.Delete(ctx); err != nil {
		return fmt.Errorf("Bucket(%q).Delete:", bucketName, err)
	}

	return nil
}

func extractFolderAndImageName(googleStorage, bucketName, url string) (string, error) {
	prefix := googleStorage + bucketName + "/"

	// Check if the URL starts with the specified prefix
	if !strings.HasPrefix(url, prefix) {
		return "", fmt.Errorf("URL does not have the expected prefix")
	}

	// Extract the folder and image name after removing the prefix
	result := strings.TrimPrefix(url, prefix)

	return result, nil
}
