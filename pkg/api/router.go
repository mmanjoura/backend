package api

import (
	"time"

	"github.com/mmanjoura/niya-voyage/backend/cmd/server/docs"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/activities"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/amadeus/flight_booking"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/comments"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/contacts"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/contentTypes"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/faqs"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/golfs"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/itineraries"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/storage"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/tours"
	"github.com/mmanjoura/niya-voyage/backend/pkg/api/users"

	"github.com/mmanjoura/niya-voyage/backend/pkg/auth"
	"github.com/mmanjoura/niya-voyage/backend/pkg/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/time/rate"
)

// InitRouter initializes the routes for the API
func InitRouter() *gin.Engine {

	r := gin.Default()

	r.Use(gin.Logger())
	// if gin.Mode() == gin.ReleaseMode {
	// 	r.Use(middleware.Security())
	// }
	r.Use(middleware.Cors())

	r.Use(middleware.RateLimiter(rate.Every(1*time.Minute), 600)) // 60 requests per minute

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{

		// tours routes
		v1.GET("/tours", tours.GetAll)
		v1.POST("/tours", middleware.JWTAuth(), tours.Create)
		v1.GET("/tours/:id", tours.GetByID)
		v1.PUT("/tours/:id", middleware.JWTAuth(), tours.UpdateByID)
		v1.DELETE("/tours/:id", middleware.JWTAuth(), tours.DeleteByID)

		// Golfs routes
		v1.GET("/golfs", golfs.GetAll)
		v1.POST("/golfs", middleware.JWTAuth(), golfs.Create)
		v1.GET("/golfs/:id", golfs.GetByID)
		v1.PUT("/golfs/:id", middleware.JWTAuth(), golfs.UpdateByID)
		v1.DELETE("/golfs/:id", middleware.JWTAuth(), golfs.DeleteByID)

		// itineraries routes
		v1.GET("/itineraries", itineraries.GetAll)
		v1.POST("/itineraries", itineraries.Create)
		v1.GET("/itineraries/:id", itineraries.GetByID)
		v1.PUT("/itineraries", itineraries.UpdateByID)
		v1.DELETE("/itineraries/:id", middleware.JWTAuth(), itineraries.DeleteByID)

		// faqs routes
		v1.GET("/faqs", faqs.GetAll)
		v1.POST("/faqs", middleware.JWTAuth(), faqs.Create)
		v1.GET("/faqs/:id", itineraries.GetByID)
		v1.PUT("/faqs", middleware.JWTAuth(), faqs.UpdateByID)
		v1.DELETE("/faqs/:id", middleware.JWTAuth(), faqs.DeleteByID)

		v1.POST("/uploadImage", storage.UploadImagesHandler)

		// Activities routes
		v1.GET("/activities", activities.GetAll)
		v1.POST("/activities", middleware.JWTAuth(), activities.Create)
		v1.GET("/activities/:id", activities.GetByID)
		v1.PUT("/activities/:id", middleware.JWTAuth(), activities.UpdateByID)
		v1.DELETE("/activities/:id", middleware.JWTAuth(), activities.DeleteByID)

		// Activities routes
		v1.GET("/activitytypes", contentTypes.GetAll)
		v1.POST("/activitytypes", middleware.JWTAuth(), contentTypes.Create)
		v1.GET("/activitytypes/:id", contentTypes.GetByID)
		v1.PUT("/activitytypes/:id", middleware.JWTAuth(), contentTypes.UpdateByID)
		v1.DELETE("/activitytypes/:id", middleware.JWTAuth(), contentTypes.DeleteByID)

		// contacts routes
		v1.POST("/contacts", contacts.Create)

		v1.POST("/comments", comments.Create)

		// amadeus routes
		v1.GET("/flightOffers", flight_booking.ShoppingFlightOffers)

		v1.POST("/login", auth.LoginHandler)
		v1.POST("/register", auth.RegisterHandler)
		v1.POST("/logout", auth.Logout)

		v1.GET("/users/account", middleware.JWTAuth(), users.Account)

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
