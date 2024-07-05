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

type UserController interface {
	// GetUser(ginCtx *gin.Context)
	GetOneUser(ginCtx *gin.Context)
	CreateUser(ginCtx *gin.Context)
	UpdateUserByEmployeeIdController(ginCtx *gin.Context)
	DeleteUserByEmployeeIdController(ginCtx *gin.Context)
}

type userController struct {
	usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) UserController {
	return &userController{
		UserUsecase: userUsecase,
	}
}

func (uc *userController) DeleteUserByEmployeeIdController(ginCtx *gin.Context) {
	ctx, cancel := context.WithTimeout(ginCtx, TIMEOUT)
	defer cancel()

	var req_id request.GetIdUserRequest

	err := ginCtx.ShouldBindUri(&req_id)
	if err != nil {
		utils.GetLogger().Error("failed validate request", zap.Error(err))
		response.BadRequest(ginCtx, err)
		return
	}

	resp, err := uc.UserUsecase.DeleteByEmployeeIdUseCase(ctx, req_id)
	if err != nil {
		response.FailedResponse(ginCtx, err)
		return
	}

	response.SuccessResponse(ginCtx, "success delete user", resp)
}

func (uc *userController) UpdateUserByEmployeeIdController(ginCtx *gin.Context) {
	ctx, cancel := context.WithTimeout(ginCtx, TIMEOUT)
	defer cancel()

	var req_id request.GetIdUserRequest

	var req request.UpdateUserRequest

	err := ginCtx.ShouldBindUri(&req_id)
	if err != nil {
		utils.GetLogger().Error("failed validate request", zap.Error(err))
		response.BadRequest(ginCtx, err)
		return
	}

	if err := request.ValidateRequest(ginCtx, &req); err != nil {
		utils.GetLogger().Error("failed validate request", zap.Error(err))
		response.BadRequest(ginCtx, err)
		return
	}

	resp, err := uc.UserUsecase.UpdateByEmployeeIdUseCase(ctx, req, req_id)
	if err != nil {
		response.FailedResponse(ginCtx, err)
		return
	}

	response.SuccessResponse(ginCtx, "success create user", resp)
}

func (uc *userController) CreateUser(ginCtx *gin.Context) {
	ctx, cancel := context.WithTimeout(ginCtx, TIMEOUT)
	defer cancel()

	var req request.CreateUserRequest

	if err := request.ValidateRequest(ginCtx, &req); err != nil {
		utils.GetLogger().Error("failed validate request", zap.Error(err))
		response.BadRequest(ginCtx, err)
		return
	}

	resp, err := uc.UserUsecase.SaveUser(ctx, req)
	if err != nil {
		response.FailedResponse(ginCtx, err)
		return
	}

	response.SuccessResponse(ginCtx, "success create user", resp)

}

// func (uc *userController) GetUser(ginCtx *gin.Context) {

// 	ctx := ginCtx.Request.Context()

// 	users, err := uc.UserUsecase.FindAllUser(ctx)
// 	if err != nil {
// 		response.FailedResponse(ginCtx, err)
// 		return
// 	}

// 	response.SuccessResponse(ginCtx, "success get user", users)
// }

func (uc *userController) GetOneUser(ginCtx *gin.Context) {
	ctx, cancel := context.WithTimeout(ginCtx, TIMEOUT)
	defer cancel()

	userIdStr := ginCtx.Param("id")

	if userIdStr == "" {
		utils.GetLogger().Error("failed validate request")
		response.BadRequest(ginCtx, nil)
		return
	}

	resp, err := uc.UserUsecase.FindUserByEmployeeId(ctx, userIdStr)

	if err != nil {
		utils.GetLogger().Error("failed validate request", zap.Error(err))
		response.FailedResponse(ginCtx, err)
		return
	}

	response.SuccessResponse(ginCtx, "success get user", resp)

}
