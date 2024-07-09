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
	GetOneProgram(ginCtx *gin.Context)
	UpdateProgramById(ginCtx *gin.Context)
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

func (controller *programController) GetOneProgram(ginCtx *gin.Context) {

	var req request.IdProgramRequest
	if err := ginCtx.ShouldBindUri(&req); err != nil {
		response.BadRequest(ginCtx, err)
		return
	}

	ctx, cancel := context.WithTimeout(ginCtx, TIMEOUT)
	defer cancel()

	resp, err := controller.ProgramUsecase.GetOneProgramUsecase(ctx, req)
	if err != nil {
		utils.GetLogger().Error("Data not found")
		response.NotFound(ginCtx, "Data not found")
		return
	}
	response.SuccessResponse(ginCtx, "success get one program", resp)
}

func (controller *programController) UpdateProgramById(ginCtx *gin.Context) {
	ctx, cancel := context.WithTimeout(ginCtx, TIMEOUT)
	defer cancel()

	var req_id request.IdProgramRequest
	if err := ginCtx.ShouldBindUri(&req_id); err != nil {
		response.BadRequest(ginCtx, err)
		return
	}

	// Check UserId
	_, err := controller.ProgramUsecase.GetOneProgramUsecase(ctx, req_id)
	if err != nil {
		response.FailedResponse(ginCtx, err)
		return
	}

	var req request.UpdateProgramRequest
	if err := request.ValidateRequest(ginCtx, &req); err != nil {
		utils.GetLogger().Error("failed validate request", zap.Error(err))
		response.BadRequest(ginCtx, err)
		return
	}

	resp, err := controller.ProgramUsecase.UpdateProgramUsecase(ctx, req, req_id)
	if err != nil {
		response.FailedResponse(ginCtx, err)
		return
	}
	response.SuccessResponse(ginCtx, "success update program", resp)
}
