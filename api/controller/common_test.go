package controller_test

import (
	"net/http/httptest"
	"template-ulamm-backend-go/api/controller"
	mock_usecase "template-ulamm-backend-go/mock/usecase"
	"template-ulamm-backend-go/pkg/model/response"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommonUsecase := mock_usecase.NewMockCommonUsecase(ctrl)

	tests := []struct {
		name       string
		mockFunc   func()
		wantErr    bool
		statusCode int
	}{
		{
			name:       "err ping",
			wantErr:    false,
			statusCode: 500,
			mockFunc: func() {
				mockCommonUsecase.EXPECT().Ping(gomock.Any()).Return(nil, assert.AnError).Times(1)
			},
		},
		{
			name:       "ping",
			wantErr:    false,
			statusCode: 200,
			mockFunc: func() {
				mockCommonUsecase.EXPECT().Ping(gomock.Any()).Return(&response.PingResponse{
					Status: "UP",
				}, nil).Times(1)
			},
		},
	}

	for _, test := range tests {
		commonController := controller.NewCommonController(mockCommonUsecase)

		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()

			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)
			commonController.Ping(ctx)

			if recorder.Code != test.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", recorder.Code, test.statusCode)
				return
			}
		})
	}
}
