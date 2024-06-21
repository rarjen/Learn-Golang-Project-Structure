package repository

import (
	"context"
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/pkg/model/entity"
	"time"
)

type CommonRepository interface {
	Ping(ctx context.Context) (*entity.Ping, error)
}

type commonRepository struct {
	Datasource *datasource.Datasource
}

func NewCommonRepository(datasource *datasource.Datasource) CommonRepository {
	return &commonRepository{
		Datasource: datasource,
	}
}

func (repo *commonRepository) Ping(ctx context.Context) (*entity.Ping, error) {
	err := repo.Datasource.PingDB()
	if err != nil {
		return nil, err
	}

	return &entity.Ping{CurrentDate: time.Now()}, nil
}
