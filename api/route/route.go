package route

import (
	"template-ulamm-backend-go/api/controller"
	"template-ulamm-backend-go/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func NewServer(controller controller.Controller) *gin.Engine {
	conf := utils.GetConfig()
	if conf.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Authorization", "Content-Type"},
	}))

	// Route Health
	routeHealth(router, controller)

	// Route Swagger
	routeSwagger(router)

	return router
}

func routeHealth(router *gin.Engine, controller controller.Controller) {
	router.GET("/health", controller.CommonController.Ping)
}

func routeSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
