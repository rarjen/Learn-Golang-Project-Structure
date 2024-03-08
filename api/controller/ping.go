package controller

import (
	"net/http"
	"template-ulamm-backend-go/pkg/datasource"

	"github.com/gin-gonic/gin"
)

type CommonController interface {
	PingDB(ginCtx *gin.Context)
}

type commonController struct {
	ds *datasource.Datasource
}

func NewCommonController(ds *datasource.Datasource) CommonController {
	return &commonController{
		ds: ds,
	}
}

func (cC *commonController) PingDB(
	ginCtx *gin.Context,
) {

	res := struct {
		Status string `json:"status"`
	}{
		Status: "UP",
	}

	if pinger, ok := cC.ds.Db.ConnPool.(interface{ Ping() error }); ok {
		if err := pinger.Ping(); err != nil {
			res.Status = err.Error()
			ginCtx.JSON(http.StatusInternalServerError, res)
		} else {
			ginCtx.JSON(http.StatusOK, res)
		}
	} else {
		res.Status = "failed to ping"
		ginCtx.JSON(http.StatusInternalServerError, res)
	}
}
