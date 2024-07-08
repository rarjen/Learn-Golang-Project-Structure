package repository

import (
	"context"
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/pkg/model/entity"
)

type ProgramRepository interface {
	CreateProgramRepo(ctx context.Context, data *entity.Program) (*entity.Program, error)
	GetAllProgramsRepo(ctx context.Context) ([]entity.Program, error)
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
		return programs, err
	}
	return programs, nil
}
