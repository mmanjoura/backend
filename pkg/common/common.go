package common

import (
	"fmt"
	"strconv"

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
