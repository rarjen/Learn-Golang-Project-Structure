package request

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"template-ulamm-backend-go/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

const HTTP_MESSAGE_INVALID_INPUT = "%s tidak boleh null & kosong"

var (
	validateCache *validator.Validate
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// Validate Incoming Request Query Parameters
func ValidateRequest(ctx *gin.Context, requestBody any) error {
	if err := ctx.ShouldBind(&requestBody); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			if len(ve) > 0 {
				newErr := fmt.Errorf(HTTP_MESSAGE_INVALID_INPUT, toSnakeCase(ve[0].Field()))
				utils.GetLogger().Error(newErr.Error())
				return newErr
			}
		}

		utils.GetLogger().Error("validate request body error", zap.Error(err))
		return err
	}

	if err := getValidator().Struct(requestBody); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			if len(ve) > 0 {
				newErr := fmt.Errorf(HTTP_MESSAGE_INVALID_INPUT, toSnakeCase(ve[0].Field()))
				utils.GetLogger().Error(newErr.Error())
				return newErr
			}
		}

		utils.GetLogger().Error("validate struct body error", zap.Error(err))
		return err
	}

	return nil
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

func getValidator() *validator.Validate {
	if validateCache == nil {
		validateCache = validator.New()
	}

	return validateCache
}

func GetLanguage(ctx *gin.Context) string {
	lang := ctx.GetHeader("Accept-Language")
	if lang == "" {
		lang = utils.DEFAULT_LANGUAGE.String()
	}

	return lang
}
