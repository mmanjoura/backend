package common

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"strconv"

	"github.com/disintegration/imaging"

	"github.com/gin-gonic/gin"
)

const (
	noTourType = iota
	Flights
	Tours
	Hotels
	Rentals
	Cars
	Golfs

	Activities
)
const (
	// silide Image
	noImage = iota
	SlideImage
	GalleryImage
)

func GetPaginationParams(c *gin.Context) (offset int, limit int) {
	offset, _ = strconv.Atoi(c.Query("offset"))
	limit, _ = strconv.Atoi(c.Query("limit"))

	offset = 0
	limit = 10

	return offset, limit
}

func BuildCondition(referrerId, categoryId int) string {
	var condition string

	if referrerId > 0 && categoryId > 0 {
		condition = fmt.Sprintf("referrer_id = %d AND category_id = %d", referrerId, categoryId)
	} else {
		condition = "1 = 1"
	}

	return condition
}

func GetCategoryId(productType string) (int, error) {
	var categoryId int

	switch productType {
	case "TOUR":
		categoryId = Tours
	case "FLIGHT":
		categoryId = Flights
	case "HOTEL":
		categoryId = Hotels
	case "RENTAL":
		categoryId = Rentals
	case "CAR":
		categoryId = Cars
	case "GOLF":
		categoryId = Golfs
	case "ACTIVITY":
		categoryId = Activities
	default:
		return 0, fmt.Errorf("Invalid product type")
	}

	return categoryId, nil
}

// ProcessImage takes an image name, crop sizes, and resize dimensions as parameters
// and performs image processing operations.
func ProcessImage(
	img image.Image,
	width,
	hight int) (image.Image, error) {
	// Open the input image file

	// Crop the original image to specified size using the center anchor.
	croppedImage := imaging.CropAnchor(img, 800, 600, imaging.Center)

	// Resize the cropped image to gallery dimensions preserving the aspect ratio.
	img = imaging.Resize(croppedImage, width, hight, imaging.Lanczos)

	return img, nil
}

func FileHeaderToReader(fileHeader *multipart.FileHeader) (io.Reader, error) {
	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}

	// Return the file reader
	return file, nil
}

func ImageToReader(img image.Image) (io.Reader, error) {
	// Create a byte buffer to store the encoded image
	var buf bytes.Buffer

	// Encode the image to JPEG format and write it to the buffer
	err := jpeg.Encode(&buf, img, nil)
	if err != nil {
		return nil, err
	}

	// Create a reader from the byte buffer
	reader := bytes.NewReader(buf.Bytes())

	return reader, nil
}
