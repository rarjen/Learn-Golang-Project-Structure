package controller

import (
	"context"
	"template-ulamm-backend-go/pkg/model/response"
	"template-ulamm-backend-go/pkg/usecase"

	"github.com/gin-gonic/gin"
)

type CommonController interface {
	Ping(ginCtx *gin.Context)
}

type commonController struct {
	usecase.CommonUsecase
}

func NewCommonController(commonUsecase usecase.CommonUsecase) CommonController {
	return &commonController{
		CommonUsecase: commonUsecase,
	}
}

// Health godoc
//
//	@Summary			Ping
//	@Description	Melakukan ping ke database untuk memeriksa kesehatan aplikasi dan database
//	@ID						get-ping
//	@Tags					ping
//	@Produce			json
//	@Success			200	{object} response.PingResponse
//	@Failure			500
//	@Router				/health [get]
func (controller *commonController) Ping(ginCtx *gin.Context) {
	ctx, cancel := context.WithTimeout(ginCtx, TIMEOUT)
	defer cancel()

	resp, err := controller.CommonUsecase.Ping(ctx)
	if err != nil {
		response.FailedResponse(ginCtx, err)
		return
	}

	response.SuccessResponse(ginCtx, "Ping Berhasil", resp)
}
