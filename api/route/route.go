package route

import (
	"fmt"
	"template-ulamm-backend-go/api/controller"
	"template-ulamm-backend-go/api/middleware"
	"template-ulamm-backend-go/utils/config"
	"template-ulamm-backend-go/utils/constantvar"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

/*
	All URI will using kebab-case
	example: http://localhost:8080/sync-data/dummy-data
*/

func NewGinServer(
	controller controller.Controller,
	conf config.Config,
) *gin.Engine {
	if conf.STAGE == constantvar.STAGE_PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	ginEngine := gin.Default()
	ginEngine.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Authorization", "Content-Type"},
	}))

	// All Public APIs
	publicRoute(ginEngine, controller)

	// All Private APIs
	privateRoute(ginEngine, controller, conf)
	return ginEngine
}

func publicRoute(
	router *gin.Engine,
	controller controller.Controller,
) {
	rgPublic := router.Group("/")
	rgPublic.GET("/health", controller.CommonController.PingDB)
	rgPublic.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func privateRoute(
	router *gin.Engine,
	controller controller.Controller,
	conf config.Config,
) {
	rgPrivate := router.Group(fmt.Sprintf("/%s", constantvar.ROUTE_SYNC_DATA))
	if conf.STAGE == constantvar.STAGE_PRODUCTION {
		rgPrivate.Use(middleware.ValidateMiddleware())
	}
	rgPrivate.GET(fmt.Sprintf("/%s", constantvar.ROUTE_SYNC_DUMMY), controller.SyncDataController.SyncDummy)
	rgPrivate.GET(fmt.Sprintf("/%s", constantvar.ROUTE_CHECK_NASABAH_ID), controller.SyncDataController.CheckDebiturID)
	rgPrivate.GET(fmt.Sprintf("/%s", constantvar.ROUTE_UNIT_BRANCH_CODE), controller.SyncDataController.GetUnitBranchCode)
}
