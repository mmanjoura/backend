package main

import (
	"log"

	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/api"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/cache"
	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/database"

	"github.com/gin-gonic/gin"
)

// @title           Niya Voyage API
// @version         1.0
// @description     This is the API documentation for the Niya Voyage API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  conatct@niyavoyage.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey JwtAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	database.ConnectDatabase()
	config := database.Database.Config
	cache.InitRedis(config["REDIS_URI"])

	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	r := api.InitRouter()

	if err := r.Run(config["PORT"]); err != nil {
		log.Fatal(err)
	}
}
