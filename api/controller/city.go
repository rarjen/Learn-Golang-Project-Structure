package controller

import (
	"net/http"
	"template-ulamm-backend-go/pkg/usecase"

	"github.com/gin-gonic/gin"
)

type CityController interface {
	GetCity(ginCtx *gin.Context)
}

type cityController struct {
	usecase.CityUsecase
}

func NewCityController(cityUsecase usecase.CityUsecase) CityController {
	return &cityController{
		CityUsecase: cityUsecase,
	}
}

func (c *cityController) GetCity(ginCtx *gin.Context) {

	cities, err := c.CityUsecase.FindAllUser()
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ginCtx.JSON(http.StatusOK, cities)

}
