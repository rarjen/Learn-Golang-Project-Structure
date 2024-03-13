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
	SyncDummy(ctx *gin.Context)
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

// SyncDummy godoc
//
//	@Summary			Get scoring parameter
//	@Description	Mendapatkan detail parameter untuk melakukan scoring
//	@ID						sync-dummy
//	@Tags					SyncDummy
//	@Produce			json
//	@Success			200	{object}	domain.SyncDataScoringParameterAPIResponse
//	@Failure			500
//	@Router				/sync-data/sync-dummy [get]
func (sDC *syncDataController) SyncDummy(
	ctx *gin.Context,
) {

	response := domain.SyncDataScoringParameterAPIResponse{
		Status:      http.StatusBadRequest,
		Description: "failed to fetch",
		Data:        nil,
	}

	data, err := sDC.syncDataUseCase.SyncDummy(ctx)
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

// CheckDebiturID godoc
//
//	@Summary			Get debitur's detail by nasabahID
//	@Description	Mendapatkan detail dari nasabah berdasarkan nasabahID
//	@ID						check-nasabah-id
//	@Tags					CheckDebiturID
//	@Param 				request body domain.CheckDebiturIDRequest true "input JSON sebagai request body"
//	@Produce			json
//	@Success			200	{object}	domain.CheckDebiturIDResponse
//	@Failure			500
//	@Router				/sync-data/check-debitur-id [get]
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

// KodeUnitCabang godoc
//
//	@Summary			Get debitur's detail by nasabahID
//	@Description	Mendapatkan detail cabang berdasarkan limit dan kriteria
//	@ID						kode-unit-cabang
//	@Tags					KodeUnitCabang
//	@Param 				request body domain.GetUnitBranchCodeRequest true "input JSON sebagai request body"
//	@Produce			json
//	@Success			200	{object}	domain.GetUnitBranchCodeResponse
//	@Failure			500
//	@Router				/sync-data/kode-unit-cabang [get]
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
