package controller

import (
	"errors"
	"template-ulamm-backend-go/pkg/datasource"
)

type CommonController interface {
	PingDB() error
}

type commonController struct {
	ds *datasource.Datasource
}

func NewCommonController(ds *datasource.Datasource) CommonController {
	return &commonController{
		ds: ds,
	}
}

func (cC *commonController) PingDB() error {
	if pinger, ok := cC.ds.Db.ConnPool.(interface{ Ping() error }); ok {
		if err := pinger.Ping(); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return errors.New("failed to ping")
	}
}
