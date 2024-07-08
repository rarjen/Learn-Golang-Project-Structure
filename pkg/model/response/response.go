package response

import (
	"net/http"
	"template-ulamm-backend-go/pkg/errs"
	"template-ulamm-backend-go/pkg/model/request"
	"time"

	"github.com/gin-gonic/gin"
)

type Error struct {
	ErrorCode string `json:"error_code"`
	Title     string `json:"title"`
	Message   string `json:"message"`
}

type Response struct {
	Error   *Error  `json:"error"`
	Message *string `json:"message"`
	Data    any     `json:"data"`
}

type CreatedUserResponse struct {
	ID         string `json:"id"`
	IDEmployee string `json:"id_employee"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	IsActive   int    `json:"is_active"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}

type CreatedProgramResponse struct {
	IDProgram   int       `json:"id"`
	ProgramName string    `json:"program_name"`
	IsActive    int       `json:"is_active"`
	CreatedBy   string    `json:"created_by"`
	ModifiedBy  string    `json:"modified_by"`
	CreatedTime time.Time `json:"created_time"`
}

type UpdatedUserResponse struct {
	ID         string `json:"id"`
	IDEmployee string `json:"id_employee"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	IsActive   int    `json:"is_active"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}

func SuccessResponse(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Message: &message,
		Data:    data,
	})
}

func BadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, Response{
		Error: &Error{
			ErrorCode: "bad_request",
			Title:     "bad request",
			Message:   err.Error(),
		},
		Data: nil,
	})
}

func FailedResponse(ctx *gin.Context, err error) {
	if serr, ok := err.(*errs.Error); ok {
		unprocessable(ctx, serr)
		return
	}

	internalServerError(ctx)
}

func internalServerError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, Response{
		// Error:   &internalServer,
		Error: &Error{
			ErrorCode: "internal_server_error",
			Title:     "internal server error",
			Message:   "internal server error",
		},
		Data: nil,
	})
}

func unprocessable(ctx *gin.Context, err *errs.Error) {
	err = errs.ErrorTranslation(err.ErrorCode(), request.GetLanguage(ctx))

	ctx.JSON(http.StatusUnprocessableEntity, Response{
		Error: &Error{
			ErrorCode: err.ErrorCode(),
			Title:     err.Title(),
			Message:   err.Message(),
		},
		Data: nil,
	})
}
