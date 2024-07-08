package usecase

import (
	"context"
	"template-ulamm-backend-go/pkg/errs"
	"template-ulamm-backend-go/pkg/model/entity"
	"template-ulamm-backend-go/pkg/model/request"
	"template-ulamm-backend-go/pkg/model/response"
	"template-ulamm-backend-go/pkg/repository"
	"template-ulamm-backend-go/utils"
	"time"

	"go.uber.org/zap"
)

type ProgramUsecase interface {
	CreateProgramUseCase(ctx context.Context, req request.CreateProgramRequest) (*response.CreatedProgramResponse, error)
}

type programUsecase struct {
	repository repository.ProgramRepository
}

func NewProgramUsecase(programRepository repository.ProgramRepository) *programUsecase {
	return &programUsecase{programRepository}
}

func (u *programUsecase) CreateProgramUseCase(ctx context.Context, req request.CreateProgramRequest) (*response.CreatedProgramResponse, error) {
	program := entity.Program{}

	program.ProgramName = req.ProgramName
	program.IsActive = req.IsActive
	program.CreatedBy = req.CreatedBy
	program.CreatedTime = time.Now()

	newProgram, err := u.repository.CreateProgramRepo(ctx, &program)
	if err != nil {
		utils.GetLogger().Error("error create program", zap.Error(err))
		return nil, errs.ERR_CREATE_PROGRAM
	}

	return &response.CreatedProgramResponse{
		IDProgram:   newProgram.IDProgram,
		ProgramName: newProgram.ProgramName,
		IsActive:    newProgram.IsActive,
		CreatedBy:   newProgram.CreatedBy,
		CreatedTime: newProgram.CreatedTime,
	}, nil

}
