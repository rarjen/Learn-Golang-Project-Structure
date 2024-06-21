package usecase

import (
	"context"
	"template-ulamm-backend-go/pkg/errs"
	"template-ulamm-backend-go/pkg/model/response"
	"template-ulamm-backend-go/pkg/repository"
	"template-ulamm-backend-go/utils"

	"go.uber.org/zap"
)

type CommonUsecase interface {
	Ping(ctx context.Context) (*response.PingResponse, error)
}

type commonUsecase struct {
	repository.CommonRepository
}

func NewCommonUsecase(commonRepository repository.CommonRepository) CommonUsecase {
	return &commonUsecase{
		CommonRepository: commonRepository,
	}
}

func (usecase *commonUsecase) Ping(ctx context.Context) (*response.PingResponse, error) {
	ping, err := usecase.CommonRepository.Ping(ctx)
	if err != nil {
		utils.GetLogger().Error("error ping database", zap.Error(err))
		return nil, errs.ERR_PING
	}

	return &response.PingResponse{
		Status:      "UP",
		CurrentDate: ping.CurrentDate,
	}, nil
}
