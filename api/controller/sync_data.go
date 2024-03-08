package controller

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"template-ulamm-backend-go/domain"
	"template-ulamm-backend-go/pkg/usecase"
	"template-ulamm-backend-go/utils/constantvar"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SyncDataController interface {
	GetDummyData() error
	ScoringUsers(ctx *gin.Context)
	CheckDebiturID(ginCtx *gin.Context)
	GetUnitBranchCode(ginCtx *gin.Context)
}

type syncDataController struct {
	syncDataUseCase usecase.SyncDataUseCase
}

func NewSyncDataController(syncDataUseCase usecase.SyncDataUseCase) SyncDataController {
	return &syncDataController{
		syncDataUseCase: syncDataUseCase,
	}
}

func (sDC *syncDataController) GetDummyData() error {
	return nil
}

func (sDC *syncDataController) ScoringUsers(
	ctx *gin.Context,
) {

	response := domain.SyncDataScoringParameterAPIResponse{
		Status:      http.StatusBadRequest,
		Description: "failed to fetch",
		Data:        nil,
	}

	data, err := sDC.syncDataUseCase.ScoringUsers(ctx)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Description = err.Error()
		ctx.JSON(response.Status, response)
		return
	}

	response.Data = data
	response.Status = http.StatusOK
	response.Description = "success"
	ctx.JSON(response.Status, response)
}

func (sDC *syncDataController) CheckDebiturID(
	ginCtx *gin.Context,
) {
	bodyReq, err := io.ReadAll(ginCtx.Request.Body)
	response := domain.CheckDebiturIDResponse{
		Status:      http.StatusBadRequest,
		Description: constantvar.HTTP_RESPONSE_FAILED_TO_FETCH,
		Data:        nil,
	}
	if err != nil {
		response.Description = err.Error()
		response.Status = http.StatusInternalServerError
		ginCtx.JSON(response.Status, response)
		return
	}
	ginCtx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyReq))

	var checkDebiturIDRequest domain.CheckDebiturIDRequest

	if err := ginCtx.ShouldBindJSON(&checkDebiturIDRequest); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			if len(ve) > 0 {
				// binding:required -> can't be empty and can't be null
				response.Description = fmt.Sprintf(constantvar.HTTP_MESSAGE_INVALID_INPUT, ve[0].Field())
				response.Status = http.StatusBadRequest
				ginCtx.JSON(response.Status, response)
				return
			}
		}
		response.Description = err.Error()
		response.Status = http.StatusInternalServerError
		ginCtx.JSON(response.Status, response)
		return
	}

	result, err := sDC.syncDataUseCase.CheckDebiturIDUnit(ginCtx, checkDebiturIDRequest)
	if err != nil {
		response.Description = err.Error()
		response.Status = http.StatusInternalServerError
		ginCtx.JSON(response.Status, response)
		return
	}

	response.Description = constantvar.HTTP_RESPONSE_SUCCESS
	response.Status = http.StatusOK
	response.Data = result

	ginCtx.JSON(response.Status, response)
}

func (sDC *syncDataController) GetUnitBranchCode(ginCtx *gin.Context) {
	response := domain.GetUnitBranchCodeResponse{
		Status:      http.StatusBadRequest,
		Description: constantvar.HTTP_RESPONSE_FAILED_TO_FETCH,
		Data:        nil,
	}
	bodyReq, err := io.ReadAll(ginCtx.Request.Body)
	if err != nil {
		ginCtx.JSON(response.Status, response)
		return
	}

	ginCtx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyReq))

	var getUnitBranchCodeRequest domain.GetUnitBranchCodeRequest
	if err := ginCtx.ShouldBindJSON(&getUnitBranchCodeRequest); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			if len(ve) > 0 {
				response.Description = err.Error()
				response.Status = http.StatusInternalServerError
				ginCtx.JSON(response.Status, response)
				return
			}
		}
		response.Description = err.Error()
		response.Status = http.StatusInternalServerError
		ginCtx.JSON(response.Status, response)
		return
	}

	data, err := sDC.syncDataUseCase.GetUnitBranchCode(ginCtx, getUnitBranchCodeRequest)
	if err != nil {
		response.Description = err.Error()
		response.Status = http.StatusInternalServerError
		ginCtx.JSON(response.Status, response)
		return
	}

	response.Data = data
	response.Description = constantvar.HTTP_RESPONSE_SUCCESS
	response.Status = http.StatusOK
	ginCtx.JSON(response.Status, response)
}
