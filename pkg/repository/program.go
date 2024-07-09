package repository

import (
	"context"
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/pkg/model/entity"
)

type ProgramRepository interface {
	CreateProgramRepo(ctx context.Context, data *entity.Program) (*entity.Program, error)
	GetAllProgramsRepo(ctx context.Context) ([]entity.Program, error)
	GetOneProgramById(ctx context.Context, id int) (entity.Program, error)
	UpdateProgramRepo(ctx context.Context, data *entity.Program) (*entity.Program, error)
	DeleteProgramRepo(ctx context.Context, data *entity.Program) error
}

type programRepository struct {
	Datasource *datasource.Datasource
}

func NewProgramRepository(datasource *datasource.Datasource) ProgramRepository {
	return &programRepository{
		Datasource: datasource,
	}
}

func (repo *programRepository) CreateProgramRepo(ctx context.Context, data *entity.Program) (*entity.Program, error) {
	err := repo.Datasource.GormDB.WithContext(ctx).Omit("id_program").Create(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *programRepository) GetAllProgramsRepo(ctx context.Context) ([]entity.Program, error) {
	var programs []entity.Program
	err := repo.Datasource.GormDB.WithContext(ctx).Order("id_program desc").Find(&programs).Error
	if err != nil {
		return nil, nil
	}
	return programs, nil
}

func (repo *programRepository) GetOneProgramById(ctx context.Context, id int) (entity.Program, error) {
	var program entity.Program
	err := repo.Datasource.GormDB.WithContext(ctx).Where("id_program = ?", id).Find(&program).Error
	if err != nil {
		return program, err
	}
	return program, nil
}

func (repo *programRepository) UpdateProgramRepo(ctx context.Context, data *entity.Program) (*entity.Program, error) {

	err := repo.Datasource.GormDB.WithContext(ctx).Model(&entity.Program{}).Where("id_program = ?", data.IDProgram).Updates(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func (repo *programRepository) DeleteProgramRepo(ctx context.Context, data *entity.Program) error {
	return repo.Datasource.GormDB.WithContext(ctx).Delete(&data, data.IDProgram).Error
}
