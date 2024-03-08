package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"template-ulamm-backend-go/domain"
	"template-ulamm-backend-go/pkg/repository"
	"template-ulamm-backend-go/utils/constantvar"
	"time"

	"github.com/gin-gonic/gin"
)

type SyncDataUseCase interface {
	GetDummyData() error
	ScoringUsers(
		c context.Context,
	) ([]domain.SyncDataScoringParameterAPIData, error)
	CheckDebiturIDUnit(
		ginCtx *gin.Context,
		nasabahData domain.CheckDebiturIDRequest,
	) ([]domain.CheckDebiturIDData, error)
	GetUnitBranchCode(
		ginCtx *gin.Context,
		requestData domain.GetUnitBranchCodeRequest,
	) ([]domain.GetUnitBranchCodeData, error)
}

type syncDataUseCase struct {
	syncDataRepository repository.SyncDataRepository
	globalRepository   repository.GlobalRepository
}

func NewSyncDataUseCase(
	syncDataRepo repository.SyncDataRepository,
	globalRepo repository.GlobalRepository,
) SyncDataUseCase {
	return &syncDataUseCase{
		syncDataRepository: syncDataRepo,
		globalRepository:   globalRepo,
	}
}

func (sDUC *syncDataUseCase) GetDummyData() error {
	return nil
}

func (sDUC *syncDataUseCase) ScoringUsers(
	c context.Context,
) ([]domain.SyncDataScoringParameterAPIData, error) {
	// create new context to handle timeout
	ctx, cancel := context.WithTimeout(c, 10*time.Second)

	defer cancel()

	response := make([]domain.SyncDataScoringParameterAPIData, 0)

	dataUsers, err := sDUC.syncDataRepository.ScoringUsers(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range dataUsers {
		response = append(response, domain.SyncDataScoringParameterAPIData{
			ID:          v.ID,
			BottomLimit: v.BottomLimit,
			UpperLimit:  v.UpperLimit,
			Category:    v.Category,
		})
	}

	return response, nil
}

func (sDUC *syncDataUseCase) CheckDebiturIDUnit(
	ginCtx *gin.Context,
	nasabahData domain.CheckDebiturIDRequest,
) ([]domain.CheckDebiturIDData, error) {

	// create new context to handle timeout
	ctx, cancel := context.WithTimeout(ginCtx, 10*time.Second)

	defer cancel()

	resultData := make([]domain.CheckDebiturIDData, 0)

	// get string connection
	newDBConnection, err := sDUC.syncDataRepository.SyncDataGetConnectionString(
		nasabahData.KodeUnit,
	)
	if err != nil {
		return nil, err
	}

	// Decode the encoded-password using API
	encodedDBpasswordRequest := domain.DecodeDBPasswordRequest{
		Unit:     newDBConnection.BranchCode,
		Password: newDBConnection.BranchPassword,
	}

	decodedBodyPassword, err := sDUC.httpRequestToDecodeDBPassword(encodedDBpasswordRequest)
	if err != nil {
		return nil, err
	}

	newDBStringConnection := fmt.Sprintf(
		constantvar.DB_SQL_STRING_CONNECTION_FORMAT,
		newDBConnection.BranchIP,
		newDBConnection.BranchUser,
		*decodedBodyPassword.DecryptedString,
		newDBConnection.BranchDatabase,
	)

	resultDB, err := sDUC.syncDataRepository.CheckDebiturIDUnit(ctx, newDBStringConnection, nasabahData)
	if err != nil {
		return nil, err
	}

	for _, v := range resultDB {
		resultData = append(resultData, domain.CheckDebiturIDData{
			BirthDate: v.BirthDate,
			NumberID:  v.NumberID,
		})
	}

	return resultData, nil
}

func (sDUC *syncDataUseCase) GetUnitBranchCode(
	ginCtx *gin.Context,
	requestData domain.GetUnitBranchCodeRequest,
) ([]domain.GetUnitBranchCodeData, error) {
	ctx, cancel := context.WithTimeout(ginCtx, 10*time.Second)
	defer cancel()

	responseData := make([]domain.GetUnitBranchCodeData, 0)

	resultDB, err := sDUC.syncDataRepository.GetUnitBranchCode(ctx, requestData)
	if err != nil {
		return nil, err
	}

	for _, v := range resultDB {
		responseData = append(responseData, domain.GetUnitBranchCodeData{
			BranchCode:   v.BranchCode,
			BranchName:   v.BranchName,
			BranchIP:     v.BranchIP,
			BranchDBName: v.BranchDBName,
		})
	}

	return responseData, nil
}

func (aUC *syncDataUseCase) httpRequestToDecodeDBPassword(
	encodedDBpasswordRequest domain.DecodeDBPasswordRequest) (
	*domain.DecodeDBPasswordResponse, error,
) {

	encodedJSON, err := json.Marshal(encodedDBpasswordRequest)
	if err != nil {
		return nil, err
	}
	apiRes, err := aUC.globalRepository.FetchAPIWithBody(
		fmt.Sprintf(
			"http://%s/%s/%s",
			constantvar.EXTERNAL_URL_MMS_FE,
			constantvar.API_LOWER_CASE,
			constantvar.ENCODED_DECODED_DB_API,
		),
		http.MethodPut,
		encodedJSON,
	)

	if err != nil {
		return nil, err
	}

	var res domain.DecodeDBPasswordResponse
	byteBody, err := io.ReadAll(apiRes.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(byteBody, &res); err != nil {
		return nil, err
	}
	if res == (domain.DecodeDBPasswordResponse{}) || res.DecryptedString == nil {
		return nil, errors.New(constantvar.HTTP_RESPONSE_DATA_NOT_FOUND)
	}
	return &res, nil
}
