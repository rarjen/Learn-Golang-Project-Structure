package middleware

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	router.Use(logMetric(
		blacklist{
			bodyRequest: []string{"password"},
		},
	))
	router.Use(corsHandler())
	router.Use(platformHandler())
}
