package controller

import (
	"context"
	"template-ulamm-backend-go/pkg/model/request"
	"template-ulamm-backend-go/pkg/model/response"
	"template-ulamm-backend-go/pkg/usecase"
	"template-ulamm-backend-go/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProgramController interface {
	CreateProgram(ginCtx *gin.Context)
	GetAllPrograms(ginCtx *gin.Context)
}

type programController struct {
	usecase.ProgramUsecase
}

func NewProgramController(programUsecase usecase.ProgramUsecase) ProgramController {
	return &programController{
		ProgramUsecase: programUsecase,
	}
}

func (controller *programController) CreateProgram(ginCtx *gin.Context) {
	ctx, cancel := context.WithTimeout(ginCtx, TIMEOUT)
	defer cancel()

	var req request.CreateProgramRequest

	if err := request.ValidateRequest(ginCtx, &req); err != nil {
		utils.GetLogger().Error("failed validate request", zap.Error(err))
		response.BadRequest(ginCtx, err)
		return
	}

	resp, err := controller.ProgramUsecase.CreateProgramUseCase(ctx, req)
	if err != nil {
		response.FailedResponse(ginCtx, err)
		return
	}
	response.SuccessResponse(ginCtx, "success create program", resp)
}

func (controller *programController) GetAllPrograms(ginCtx *gin.Context) {
	ctx, cancel := context.WithTimeout(ginCtx, TIMEOUT)
	defer cancel()

	resp, err := controller.ProgramUsecase.GetAllProgramsUsecase(ctx)
	if err != nil {
		response.FailedResponse(ginCtx, err)
		return
	}
	response.SuccessResponse(ginCtx, "success get all programs", resp)
}
