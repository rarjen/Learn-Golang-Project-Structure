package repository

import (
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/pkg/model/entity"
)

type PipelineRepository interface {
	// Save(pipeline *entity.Pipeline) error
}

type pipelineRepository struct {
	Datasource *datasource.Datasource
}

func NewPipelineRepository(datasource *datasource.Datasource) PipelineRepository {
	return &pipelineRepository{
		Datasource: datasource,
	}
}

func (cr *pipelineRepository) FindAll() ([]entity.User, error) {
	var users []entity.User

	err := cr.Datasource.GormDB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
