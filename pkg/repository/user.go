package repository

import (
	"context"
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/pkg/model/entity"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	FindById(ctx context.Context, id string) (entity.User, error)
	Save(context.Context, *entity.User) error
}

type userRepository struct {
	Datasource *datasource.Datasource
}

func NewUserRepository(datasource *datasource.Datasource) UserRepository {
	return &userRepository{
		Datasource: datasource,
	}
}

func (repo *userRepository) Save(ctx context.Context, data *entity.User) error {
	return repo.Datasource.GormDB.WithContext(ctx).Omit("id").Create(&data).Error
}

func (cr *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	err := cr.Datasource.GormDB.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (cr *userRepository) FindById(ctx context.Context, id string) (entity.User, error) {
	var user entity.User = entity.User{IDEmployee: id}

	err := cr.Datasource.GormDB.WithContext(ctx).Where("id_employee = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
