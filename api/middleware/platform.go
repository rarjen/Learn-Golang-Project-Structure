package middleware

import (
	"template-ulamm-backend-go/pkg/constantvar"

	"github.com/gin-gonic/gin"
)

func platformHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		platformType := ctx.GetHeader(constantvar.HEADER_PLATFORM_TYPE)

		if platformType == constantvar.HEADER_PLATFORM_TYPE_MOBILE {
			ctx.Set(constantvar.CONTEXT_PLATFORM_TYPE, constantvar.HEADER_PLATFORM_TYPE_MOBILE)
		} else {
			ctx.Set(constantvar.CONTEXT_PLATFORM_TYPE, constantvar.HEADER_PLATFORM_TYPE_WEB)
		}
	}
}
