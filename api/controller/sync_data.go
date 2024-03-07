package controller

import (
	"bytes"
	"context"
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
	ScoringUsers(
		ctx context.Context,
	) (domain.SyncDataScoringParameterAPIResponse, error)
	CheckDebiturID(
		ginCtx *gin.Context,
	) (domain.CheckDebiturIDResponse, error)
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
	ctx context.Context,
) (domain.SyncDataScoringParameterAPIResponse, error) {

	response := domain.SyncDataScoringParameterAPIResponse{
		Status:      http.StatusBadRequest,
		Description: "failed to fetch",
		Data:        nil,
	}

	data, err := sDC.syncDataUseCase.ScoringUsers(ctx)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Description = err.Error()
		return response, err
	}

	response.Data = data
	response.Status = http.StatusOK
	response.Description = "success"

	return response, nil
}

func (sDC *syncDataController) CheckDebiturID(
	ginCtx *gin.Context,
) (domain.CheckDebiturIDResponse, error) {
	bodyReq, err := io.ReadAll(ginCtx.Request.Body)
	response := domain.CheckDebiturIDResponse{
		Status:      http.StatusBadRequest,
		Description: constantvar.HTTP_RESPONSE_FAILED_TO_FETCH,
		Data:        nil,
	}
	if err != nil {
		response.Description = err.Error()
		response.Status = http.StatusInternalServerError
		return response, err
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
				return response, err
			}
		}
		response.Description = err.Error()
		response.Status = http.StatusInternalServerError
		return response, err
	}

	result, err := sDC.syncDataUseCase.CheckDebiturIDUnit(ginCtx, checkDebiturIDRequest)
	if err != nil {
		response.Description = err.Error()
		response.Status = http.StatusInternalServerError
		return response, err
	}

	response.Description = constantvar.HTTP_RESPONSE_SUCCESS
	response.Status = http.StatusOK
	response.Data = result

	return response, nil
}
