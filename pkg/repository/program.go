package repository

import (
	"context"
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/pkg/model/entity"
)

type ProgramRepository interface {
	CreateProgramRepo(ctx context.Context, data *entity.Program) (*entity.Program, error)
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
