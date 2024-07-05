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
	GetUser(ginCtx *gin.Context)
	// GetOneUser(ginCtx *gin.Context)
	CreateUser(ginCtx *gin.Context)
}

type userController struct {
	usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) UserController {
	return &userController{
		UserUsecase: userUsecase,
	}
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

func (uc *userController) GetUser(ginCtx *gin.Context) {

	ctx := ginCtx.Request.Context()

	users, err := uc.UserUsecase.FindAllUser(ctx)
	if err != nil {
		response.FailedResponse(ginCtx, err)
		return
	}

	response.SuccessResponse(ginCtx, "success get user", users)
}

// func (uc *userController) GetOneUser(ginCtx *gin.Context) {

// 	user, err := uc.UserUsecase.FindUserById(ginCtx.Param("id"))

// 	if err != nil {
// 		response.FailedResponse(ginCtx, err)
// 		return
// 	}
// 	response.SuccessResponse(ginCtx, "success get user", user)

// }
