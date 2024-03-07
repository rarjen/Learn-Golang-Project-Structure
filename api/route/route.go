package route

import (
	"fmt"
	"net/http"
	"template-ulamm-backend-go/api/controller"
	"template-ulamm-backend-go/api/middleware"
	"template-ulamm-backend-go/utils/config"
	"template-ulamm-backend-go/utils/constantvar"

	"github.com/gin-gonic/gin"
)

/*
	All URI will using kebab-case
	example: http://localhost:8080/sync-data/dummy-data
*/

func NewGinServer(
	controller controller.Controller,
	conf config.Config,
) *gin.Engine {

	ginEngine := gin.Default()

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
	rgPublic.GET("/health", func(ctx *gin.Context) {
		if err := controller.CommonController.PingDB(); err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
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
	rgPrivate.GET(fmt.Sprintf("/%s", constantvar.ROUTE_SYNC_DUMMY), func(ctx *gin.Context) {
		res, err := controller.SyncDataController.ScoringUsers(ctx)
		if err != nil {
			ctx.JSON(res.Status, res)
			return
		}
		ctx.JSON(http.StatusOK, res)
	})
	rgPrivate.POST(fmt.Sprintf("/%s", constantvar.ROUTE_CHECK_NASABAH_ID), func(ctx *gin.Context) {
		res, err := controller.SyncDataController.CheckDebiturID(ctx)
		if err != nil {
			ctx.JSON(res.Status, res)
			return
		}
		ctx.JSON(http.StatusOK, res)
	})
}
