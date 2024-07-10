package controller

import (
	"context"
	"template-ulamm-backend-go/pkg/model/request"
	"template-ulamm-backend-go/pkg/model/response"
	"template-ulamm-backend-go/pkg/usecase"
	"template-ulamm-backend-go/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductController interface {
	GetAllProducts(ginCtx *gin.Context)
	GetOneProduct(ginCtx *gin.Context)
	CreateProduct(ginCtx *gin.Context)
	UpdateProduct(ginCtx *gin.Context)
	DeleteProduct(ginCtx *gin.Context)
}

type productController struct {
	usecase.ProductUsecase
}

func NewProductController(productUsecase usecase.ProductUsecase) ProductController {
	return &productController{
		ProductUsecase: productUsecase,
	}
}

func (pc *productController) GetAllProducts(ginCtx *gin.Context) {
	ctx, cancel := context.WithTimeout(ginCtx, TIMEOUT)
	defer cancel()
	responses, err := pc.ProductUsecase.GetAllProducts(ctx)
	if err != nil {
		utils.GetLogger().Error("failed get all programs", zap.Error(err))
		response.FailedResponse(ginCtx, err)
		return
	}
	response.SuccessResponse(ginCtx, "success get all programs", responses)
}

func (pc *productController) GetOneProduct(ginCtx *gin.Context) {

	var req request.IdProductParam
	if err := ginCtx.ShouldBindUri(&req); err != nil {
		response.BadRequest(ginCtx, err)
		return
	}

	ctx, cancel := context.WithTimeout(ginCtx, TIMEOUT)
	defer cancel()
	resp, err := pc.ProductUsecase.GetOneProduct(ctx, req.IDProduct)
	if err != nil {
		utils.GetLogger().Error("Data not found")
		response.NotFound(ginCtx, "Data not found")
		return
	}
	response.SuccessResponse(ginCtx, "success get one program", resp)
}

func (pc *productController) CreateProduct(ginCtx *gin.Context) {
	ctx, cancel := context.WithTimeout(ginCtx, TIMEOUT)
	defer cancel()

	var reqBody request.CreateProductRequest
	if err := request.ValidateRequest(ginCtx, &reqBody); err != nil {
		utils.GetLogger().Error("failed validate request", zap.Error(err))
		response.FailedResponse(ginCtx, err)
		return
	}

	resp, err := pc.ProductUsecase.CreateProduct(ctx, reqBody)
	if err != nil {
		response.FailedResponse(ginCtx, err)
		return
	}
	response.SuccessResponse(ginCtx, "success create product", resp)
}

func (pc *productController) UpdateProduct(ginCtx *gin.Context) {
	ctx, cancel := context.WithTimeout(ginCtx, TIMEOUT)
	defer cancel()

	var req request.IdProductParam
	if err := ginCtx.ShouldBindUri(&req); err != nil {
		response.BadRequest(ginCtx, err)
		return
	}

	productExist, err := pc.ProductUsecase.GetOneProduct(ctx, req.IDProduct)
	if err != nil {
		utils.GetLogger().Error("failed validate request", zap.Error(err))
		response.FailedResponse(ginCtx, err)
		return
	}

	if productExist.IDProduct == 0 {
		utils.GetLogger().Error("Data not found")
		response.NotFound(ginCtx, "Data not found")
		return
	}

	var reqBody request.CreateProductRequest
	if err := request.ValidateRequest(ginCtx, &reqBody); err != nil {
		utils.GetLogger().Error("failed validate request", zap.Error(err))
		response.FailedResponse(ginCtx, err)
		return
	}

	_, err = pc.ProductUsecase.UpdateProduct(ctx, reqBody, req.IDProduct)
	if err != nil {
		response.FailedResponse(ginCtx, err)
		return
	}

	updatedProduct, err := pc.ProductUsecase.GetOneProduct(ctx, req.IDProduct)
	if err != nil {
		utils.GetLogger().Error("failed validate request", zap.Error(err))
		response.FailedResponse(ginCtx, err)
		return
	}

	response.SuccessResponse(ginCtx, "success update product", updatedProduct)
}

func (pc *productController) DeleteProduct(ginCtx *gin.Context) {
	ctx, cancel := context.WithTimeout(ginCtx, TIMEOUT)
	defer cancel()
	var req request.IdProductParam
	if err := ginCtx.ShouldBindUri(&req); err != nil {
		response.BadRequest(ginCtx, err)
		return
	}

	_, err := pc.ProductUsecase.GetOneProduct(ctx, req.IDProduct)
	if err != nil {
		utils.GetLogger().Error("Data not found")
		response.NotFound(ginCtx, "Data not found")
		return
	}

	_, err = pc.ProductUsecase.DeleteProduct(ctx, req.IDProduct)
	if err != nil {
		response.FailedResponse(ginCtx, err)
		return
	}
	response.SuccessResponse(ginCtx, "success delete product", 1)
}
