package flight_booking

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/amadeus"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/proxy"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// Healthcheck godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router / [get]
func Healthcheck(g *gin.Context) {
	g.JSON(http.StatusOK, "ok")
}

// FindHotelBookings godoc
// @Summary Get all hotelBookings with pagination
// @Description Get a list of all hotelBookings with optional pagination
// @Tags hotelBookings
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.HotelBooking "Successfully retrieved list of hotelBookings"
// @Router /hotelBookings [get]
func ShoppingFlightOffers(c *gin.Context) {

	client, err := proxy.GetAmadiusClient()
	if err != nil {
		fmt.Println("not expected error while creating client", err)
	}

	// get offer flights request&response
	offerReq, offerResp, err := client.NewRequest(amadeus.ShoppingFlightOffers)

	// set offer flights params
	offerReq.(*amadeus.ShoppingFlightOffersRequest).SetCurrency("USD").SetSources("GDS").Return(
		"LAX",
		"NYC",
		time.Now().AddDate(0, 5, 0).Format("2006-01-02"),
		time.Now().AddDate(0, 7, 0).Format("2006-01-02"),
	).AddTravelers(1, 0, 0)

	// send request
	err = client.Do(offerReq, &offerResp, "GET")

	// get response
	offerRespData := offerResp.(*amadeus.ShoppingFlightOffersResponse)

	c.JSON(http.StatusOK, gin.H{"data": offerRespData})
}
