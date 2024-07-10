package repository

import (
	"context"
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/pkg/model/entity"
)

type ProductRepository interface {
	GetAllProducts(ctx context.Context) ([]*entity.Product, error)
	GetOneProduct(ctx context.Context, id int) (*entity.Product, error)
	CreateProduct(ctx context.Context, data *entity.Product) (*entity.Product, error)
	UpdateProduct(ctx context.Context, reqBody *entity.Product, id int) (*entity.Product, error)
	DeleteProduct(ctx context.Context, id int) error
}

type productRepository struct {
	Datasource *datasource.Datasource
}

func NewProductRepository(datasource *datasource.Datasource) ProductRepository {
	return &productRepository{
		Datasource: datasource,
	}
}

func (pr *productRepository) GetAllProducts(ctx context.Context) ([]*entity.Product, error) {
	var products []*entity.Product
	err := pr.Datasource.GormDB.WithContext(ctx).Order("id_product desc").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *productRepository) GetOneProduct(ctx context.Context, id int) (*entity.Product, error) {
	var product *entity.Product
	err := pr.Datasource.GormDB.WithContext(ctx).Where("id_product = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pr *productRepository) CreateProduct(ctx context.Context, data *entity.Product) (*entity.Product, error) {
	err := pr.Datasource.GormDB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (pr *productRepository) UpdateProduct(ctx context.Context, reqBody *entity.Product, id int) (*entity.Product, error) {
	err := pr.Datasource.GormDB.WithContext(ctx).Model(&entity.Product{}).Where("id_product = ?", id).Updates(&reqBody).Error
	if err != nil {
		return nil, err
	}
	return reqBody, nil
}

func (pr *productRepository) DeleteProduct(ctx context.Context, id int) error {
	err := pr.Datasource.GormDB.WithContext(ctx).Where("id_product = ?", id).Delete(&entity.Product{}).Error
	if err != nil {
		return err
	}
	return nil
}
