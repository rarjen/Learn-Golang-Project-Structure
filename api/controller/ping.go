package controller

import (
	"net/http"
	"template-ulamm-backend-go/domain"
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/utils/constantvar"

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

// Health godoc
//
//	@Summary			Ping the DB
//	@Description	Melakukan ping ke database untuk memeriksa kesehatan aplikasi dan database
//	@ID						get-pingdb
//	@Tags					pingdb
//	@Produce			json
//	@Success			200	{object} domain.PingDBResponse
//	@Failure			500
//	@Router				/health [get]
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
			ginCtx.JSON(http.StatusOK, domain.PingDBResponse{
				Status: constantvar.HTTP_RESPONSE_UP,
			})
		}
	} else {
		res.Status = "failed to ping"
		ginCtx.JSON(http.StatusInternalServerError, res)
	}
}
