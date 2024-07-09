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
	GetAllProgramsUsecase(ctx context.Context) ([]response.GetAllProgramsResponse, error)
	GetOneProgramUsecase(ctx context.Context, input request.IdProgramRequest) (*response.GetOneProgramResponse, error)
	UpdateProgramUsecase(ctx context.Context, req request.UpdateProgramRequest, req_id request.IdProgramRequest) (*response.UpdatedProgramResponse, error)
	DeleteProgramUsecase(ctx context.Context, data *entity.Program) (int, error)
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

func (u *programUsecase) GetAllProgramsUsecase(ctx context.Context) ([]response.GetAllProgramsResponse, error) {

	programs, err := u.repository.GetAllProgramsRepo(ctx)

	if err != nil {
		utils.GetLogger().Error("error get all programs", zap.Error(err))
		return nil, errs.ERR_GET_ALL_PROGRAMS
	}

	responses := make([]response.GetAllProgramsResponse, 0, len(programs))
	for _, program := range programs {
		responses = append(responses, response.GetAllProgramsResponse{
			IDProgram:   program.IDProgram,
			ProgramName: program.ProgramName,
			IsActive:    program.IsActive,
			CreatedBy:   program.CreatedBy,
			ModifiedBy:  program.ModifiedBy,
		})
	}
	return responses, nil
}

func (u *programUsecase) GetOneProgramUsecase(ctx context.Context, input request.IdProgramRequest) (*response.GetOneProgramResponse, error) {

	existingProgram, err := u.repository.GetOneProgramById(ctx, input.IDProgram)
	if err != nil {
		utils.GetLogger().Error("error get one program", zap.Error(err))
		return nil, errs.ERR_GET_ONE_PROGRAM
	}

	if existingProgram.IDProgram == 0 {
		utils.GetLogger().Error("program not found", zap.Error(err))
		return nil, errs.ERR_INIT_NOT_FOUND
	}

	return &response.GetOneProgramResponse{
		IDProgram:    existingProgram.IDProgram,
		ProgramName:  existingProgram.ProgramName,
		IsActive:     existingProgram.IsActive,
		CreatedBy:    existingProgram.CreatedBy,
		ModifiedBy:   existingProgram.ModifiedBy,
		CreatedTime:  existingProgram.CreatedTime,
		ModifiedTime: existingProgram.ModifiedTime,
	}, nil

}

func (u *programUsecase) UpdateProgramUsecase(ctx context.Context, req request.UpdateProgramRequest, req_id request.IdProgramRequest) (*response.UpdatedProgramResponse, error) {

	program := entity.Program{}

	program.IDProgram = req_id.IDProgram
	program.ProgramName = req.ProgramName
	program.IsActive = req.IsActive
	program.ModifiedBy = req.ModifiedBy
	program.ModifiedTime = time.Now()

	updatedProgram, err := u.repository.UpdateProgramRepo(ctx, &program)
	if err != nil {
		utils.GetLogger().Error("error update program", zap.Error(err))
		return nil, errs.ERR_UPDATE_PROGRAM
	}

	return &response.UpdatedProgramResponse{
		IDProgram:    updatedProgram.IDProgram,
		ProgramName:  updatedProgram.ProgramName,
		IsActive:     updatedProgram.IsActive,
		ModifiedBy:   updatedProgram.ModifiedBy,
		ModifiedTime: updatedProgram.ModifiedTime,
		CreatedTime:  updatedProgram.CreatedTime,
		CreatedBy:    updatedProgram.CreatedBy,
	}, nil
}

func (u *programUsecase) DeleteProgramUsecase(ctx context.Context, data *entity.Program) (int, error) {

	err := u.repository.DeleteProgramRepo(ctx, data)
	if err != nil {
		return 0, err
	}

	return 1, nil
}
