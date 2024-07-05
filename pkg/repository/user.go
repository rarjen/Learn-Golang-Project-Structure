package repository

import (
	"context"
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/pkg/model/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	// FindAll(ctx context.Context) ([]entity.User, error)
	FindByEmployeeId(context.Context, string) (*entity.User, error)
	Save(context.Context, *entity.User) error
	UpdateByEmployeeIdRepo(context.Context, *entity.User) error
	DeleteByEmployeeIdRepo(context.Context, string) error
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
	return repo.Datasource.GormDB.WithContext(ctx).Create(&data).Error
}

func (repo *userRepository) UpdateByEmployeeIdRepo(ctx context.Context, data *entity.User) error {
	return repo.Datasource.GormDB.WithContext(ctx).Updates(&data).Error
}

func (repo *userRepository) FindByEmployeeId(ctx context.Context, id string) (*entity.User, error) {
	var result *entity.User

	err := repo.Datasource.GormDB.WithContext(ctx).Where("id_employee = ?", id).First(&result).Error

	return result, err
}

func (repo *userRepository) DeleteByEmployeeIdRepo(ctx context.Context, id string) error {
	exec := repo.Datasource.GormDB.WithContext(ctx).Where("id_employee = ?", id).Delete(&entity.User{})
	if err := exec.Error; err != nil {
		return err
	}
	if exec.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// func (cr *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {
// 	var users []entity.User

// 	err := cr.Datasource.GormDB.WithContext(ctx).Find(&users).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }
