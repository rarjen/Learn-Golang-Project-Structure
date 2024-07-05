package usecase

import (
	"context"
	"errors"
	"fmt"
	"template-ulamm-backend-go/pkg/errs"
	"template-ulamm-backend-go/pkg/model/entity"
	"template-ulamm-backend-go/pkg/model/request"
	"template-ulamm-backend-go/pkg/model/response"
	"template-ulamm-backend-go/utils"
	"time"

	"template-ulamm-backend-go/pkg/repository"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserUsecase interface {
	// FindAllUser(ctx context.Context) ([]entity.User, error)
	FindUserByEmployeeId(ctx context.Context, id string) (*entity.User, error)
	SaveUser(ctx context.Context, req request.CreateUserRequest) (*response.CreatedUserResponse, error)
	UpdateByEmployeeIdUseCase(ctx context.Context, req request.UpdateUserRequest, req_id request.GetIdUserRequest) (*response.UpdatedUserResponse, error)
	DeleteByEmployeeIdUseCase(ctx context.Context, req_id request.GetIdUserRequest) (int, error)
}

type userUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository}
}

func (cu *userUsecase) UpdateByEmployeeIdUseCase(ctx context.Context, req request.UpdateUserRequest, req_id request.GetIdUserRequest) (*response.UpdatedUserResponse, error) {
	userExist, err := cu.repository.FindByEmployeeId(ctx, req_id.IDEmployee)
	if err != nil {
		return nil, err
	}

	userExist.IDEmployee = req.IDEmployee
	userExist.Username = req.Username
	userExist.Name = req.Name
	userExist.IsActive = req.IsActive
	userExist.CreatedBy = req.CreatedBy
	userExist.CreatedTime = time.Now()
	userExist.ModifiedBy = req.ModifiedBy
	userExist.ModifiedTime = time.Now()

	err = cu.repository.UpdateByEmployeeIdRepo(ctx, userExist)
	if err != nil {
		utils.GetLogger().Error("error update user", zap.Error(err))
		return nil, errs.ERR_UPDATE_USER
	}

	return &response.UpdatedUserResponse{
		IDEmployee: userExist.IDEmployee,
		Username:   userExist.Username,
		Name:       userExist.Name,
		IsActive:   userExist.IsActive,
		CreatedBy:  userExist.CreatedBy,
		ModifiedBy: userExist.ModifiedBy,
	}, nil
}

func (cu *userUsecase) DeleteByEmployeeIdUseCase(ctx context.Context, req_id request.GetIdUserRequest) (int, error) {

	// _, err := cu.repository.FindByEmployeeId(ctx, req_id.IDEmployee)
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		utils.GetLogger().Error("user not found")
	// 		return 0, errs.ERR_USER_NOT_FOUND
	// 	}

	// 	return 0, err
	// }

	err := cu.repository.DeleteByEmployeeIdRepo(ctx, req_id.IDEmployee)
	if err != nil {
		utils.GetLogger().Error("error delete user", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.GetLogger().Error("user not found")
			return 0, errs.ERR_USER_NOT_FOUND
		}

		return 0, errs.ERR_DELETE_USER
	}

	return 1, nil
}

// func (cu *userUsecase) FindAllUser(ctx context.Context) ([]entity.User, error) {
// 	campaigns, err := cu.repository.FindAll(ctx)
// 	if err != nil {
// 		return campaigns, err
// 	}

// 	return campaigns, nil
// }

func (cu *userUsecase) FindUserByEmployeeId(ctx context.Context, id string) (*entity.User, error) {
	userExist, err := cu.repository.FindByEmployeeId(ctx, id)
	if err != nil {
		return userExist, err
	}
	return userExist, err
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
