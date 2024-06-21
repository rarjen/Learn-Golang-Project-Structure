package route

import (
	"template-ulamm-backend-go/api/middleware"

	"github.com/gin-gonic/gin"
)

func routeMiddleware(router *gin.Engine) {
	middleware.Register(router)
}
