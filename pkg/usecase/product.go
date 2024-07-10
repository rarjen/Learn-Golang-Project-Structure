package usecase

import (
	"context"
	"fmt"
	"template-ulamm-backend-go/pkg/errs"
	"template-ulamm-backend-go/pkg/model/entity"
	"template-ulamm-backend-go/pkg/model/request"
	"template-ulamm-backend-go/pkg/model/response"
	"template-ulamm-backend-go/pkg/repository"
	"template-ulamm-backend-go/utils"
	"time"

	"go.uber.org/zap"
)

type ProductUsecase interface {
	GetAllProducts(ctx context.Context) ([]*response.GetAllProductsResponse, error)
	GetOneProduct(ctx context.Context, id int) (*response.GetAllProductsResponse, error)
	CreateProduct(ctx context.Context, data request.CreateProductRequest) (*response.GetAllProductsResponse, error)
	UpdateProduct(ctx context.Context, reqBody request.CreateProductRequest, id int) (*entity.Product, error)
	DeleteProduct(ctx context.Context, id int) (int, error)
}

type productUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repository repository.ProductRepository) *productUsecase {
	return &productUsecase{
		repository: repository,
	}
}

func (pu *productUsecase) GetAllProducts(ctx context.Context) ([]*response.GetAllProductsResponse, error) {
	product, err := pu.repository.GetAllProducts(ctx)
	if err != nil {
		utils.GetLogger().Error("failed get all programs", zap.Error(err))
		return nil, errs.ERR_GET_ALL_PROGRAMS
	}

	responses := make([]*response.GetAllProductsResponse, 0, len(product))
	for _, p := range product {
		responses = append(responses, &response.GetAllProductsResponse{
			IDProduct:          p.IDProduct,
			ProductName:        p.ProductName,
			ProductCode:        p.ProductCode,
			InterestRate:       p.InterestRate,
			InterestRateAnnual: p.InterestRateAnnual,
			LimitLoanLower:     p.LimitLoanLower,
			LimitLoanUpper:     p.LimitLoanUpper,
			TimePeriodLower:    p.TimePeriodLower,
			TimePeriodUpper:    p.TimePeriodUpper,
			IsActive:           p.IsActive,
			CreatedBy:          p.CreatedBy,
			CreatedTime:        p.CreatedTime,
			ModifiedBy:         p.ModifiedBy,
			ModifiedTime:       p.ModifiedTime,
		})
	}
	return responses, nil
}

func (pu *productUsecase) GetOneProduct(ctx context.Context, id int) (*response.GetAllProductsResponse, error) {
	product, err := pu.repository.GetOneProduct(ctx, id)
	if err != nil {
		utils.GetLogger().Error("failed get all programs", zap.Error(err))
		return nil, errs.ERR_GET_ALL_PROGRAMS
	}
	return &response.GetAllProductsResponse{
		IDProduct:          product.IDProduct,
		ProductName:        product.ProductName,
		ProductCode:        product.ProductCode,
		InterestRate:       product.InterestRate,
		InterestRateAnnual: product.InterestRateAnnual,
		LimitLoanLower:     product.LimitLoanLower,
		LimitLoanUpper:     product.LimitLoanUpper,
		TimePeriodLower:    product.TimePeriodLower,
		TimePeriodUpper:    product.TimePeriodUpper,
		IsActive:           product.IsActive,
		CreatedBy:          product.CreatedBy,
		CreatedTime:        product.CreatedTime,
		ModifiedBy:         product.ModifiedBy,
		ModifiedTime:       product.ModifiedTime,
	}, nil
}

func (pu *productUsecase) CreateProduct(ctx context.Context, data request.CreateProductRequest) (*response.GetAllProductsResponse, error) {
	product := entity.Product{}

	fmt.Println("data", data.ProductName)

	product.ProductName = data.ProductName
	product.ProductCode = data.ProductCode
	product.InterestRate = data.InterestRate
	product.InterestRateAnnual = data.InterestRateAnnual
	product.LimitLoanLower = data.LimitLoanLower
	product.LimitLoanUpper = data.LimitLoanUpper
	product.TimePeriodLower = data.TimePeriodLower
	product.TimePeriodUpper = data.TimePeriodUpper
	product.IsActive = data.IsActive
	product.CreatedBy = data.CreatedBy
	product.CreatedTime = time.Now()
	product.ModifiedTime = time.Now()

	createdProduct, err := pu.repository.CreateProduct(ctx, &product)

	if err != nil {
		utils.GetLogger().Error("failed create product", zap.Error(err))
		return nil, errs.ERR_CREATE_PRODUCT
	}
	return &response.GetAllProductsResponse{
		IDProduct:          createdProduct.IDProduct,
		ProductCode:        createdProduct.ProductCode,
		InterestRate:       createdProduct.InterestRate,
		InterestRateAnnual: createdProduct.InterestRateAnnual,
		LimitLoanLower:     createdProduct.LimitLoanLower,
		LimitLoanUpper:     createdProduct.LimitLoanUpper,
		TimePeriodLower:    createdProduct.TimePeriodLower,
		TimePeriodUpper:    createdProduct.TimePeriodUpper,
		IsActive:           createdProduct.IsActive,
		CreatedBy:          createdProduct.CreatedBy,
		CreatedTime:        createdProduct.CreatedTime,
		ModifiedBy:         createdProduct.ModifiedBy,
		ModifiedTime:       createdProduct.ModifiedTime,
	}, nil
}

func (pu *productUsecase) UpdateProduct(ctx context.Context, reqBody request.CreateProductRequest, id int) (*entity.Product, error) {

	product := entity.Product{
		ProductName:        reqBody.ProductName,
		ProductCode:        reqBody.ProductCode,
		InterestRate:       reqBody.InterestRate,
		InterestRateAnnual: reqBody.InterestRateAnnual,
		LimitLoanLower:     reqBody.LimitLoanLower,
		LimitLoanUpper:     reqBody.LimitLoanUpper,
		TimePeriodLower:    reqBody.TimePeriodLower,
		TimePeriodUpper:    reqBody.TimePeriodUpper,
		IsActive:           reqBody.IsActive,
		ModifiedBy:         reqBody.ModifiedBy,
		ModifiedTime:       time.Now(),
	}

	updatedProduct, err := pu.repository.UpdateProduct(ctx, &product, id)
	if err != nil {
		utils.GetLogger().Error("failed update product", zap.Error(err))
		return nil, errs.ERR_UPDATE_PRODUCT
	}

	return updatedProduct, nil
}

func (pu *productUsecase) DeleteProduct(ctx context.Context, id int) (int, error) {
	err := pu.repository.DeleteProduct(ctx, id)
	if err != nil {
		utils.GetLogger().Error("failed delete product", zap.Error(err))
		return 0, errs.ERR_DELETE_PRODUCT
	}
	return 1, nil
}
