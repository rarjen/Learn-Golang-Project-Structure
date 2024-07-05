package usecase

import (
	"context"
	"fmt"
	"template-ulamm-backend-go/pkg/errs"
	"template-ulamm-backend-go/pkg/model/entity"
	"template-ulamm-backend-go/pkg/model/request"
	"template-ulamm-backend-go/pkg/model/response"
	"template-ulamm-backend-go/utils"
	"time"

	"template-ulamm-backend-go/pkg/repository"

	"go.uber.org/zap"
)

type UserUsecase interface {
	FindAllUser(ctx context.Context) ([]entity.User, error)
	FindUserById(ctx context.Context, id string) (entity.User, error)
	SaveUser(ctx context.Context, req request.CreateUserRequest) (*response.CreatedUserResponse, error)
}

type userUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository}
}

func (cu *userUsecase) FindAllUser(ctx context.Context) ([]entity.User, error) {
	campaigns, err := cu.repository.FindAll(ctx)
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (cu *userUsecase) FindUserById(ctx context.Context, id string) (entity.User, error) {
	campaign, err := cu.repository.FindById(ctx, id)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (usecase *userUsecase) SaveUser(ctx context.Context, req request.CreateUserRequest) (*response.CreatedUserResponse, error) {

	user := entity.User{
		IDEmployee:   req.IDEmployee,
		Username:     req.Username,
		Name:         req.Name,
		IsActive:     req.IsActive,
		CreatedBy:    req.CreatedBy,
		CreatedTime:  time.Now(),
		ModifiedBy:   req.ModifiedBy,
		ModifiedTime: time.Now(),
	}

	fmt.Println(user)

	err := usecase.repository.Save(ctx, &user)
	if err != nil {
		utils.GetLogger().Error("error create pipeline", zap.Error(err))
		return nil, errs.ERR_CREATE_USER
	}

	return &response.CreatedUserResponse{}, nil
}
