package middleware

import (
	"errors"
	"net/http"
	"strings"
	"template-ulamm-backend-go/utils/config"
	"template-ulamm-backend-go/utils/constantvar"
	"template-ulamm-backend-go/utils/httputility"

	"github.com/gin-gonic/gin"
)

// this middleware used to validate each request.
// will forwarded the request to auth service
func ValidateMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get the header

		authToken, err := getTokenString(ctx.GetHeader(constantvar.AUTHORIZATION_SPECIAL_CASE))
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// hit the external api to forward the request
		httpRes, err := httputility.RESTValidateAccessToken(
			config.C.AuthAPIUrl+constantvar.ROUTE_API_AUTH,
			authToken,
		)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// check if status code is 202
		if httpRes.StatusCode == http.StatusAccepted {
			// continue the control to next middleware/handler
			ctx.Next()
		} else {
			// return any code from auth service
			ctx.AbortWithStatus(httpRes.StatusCode)
			return
		}
	}
}

func getTokenString(header string) (string, error) {
	if !strings.Contains(header, "Bearer") {
		return "", errors.New(constantvar.JWT_TOKEN_FAILED_INVALID_BEARER_TOKEN)
	}

	return header, nil
}
