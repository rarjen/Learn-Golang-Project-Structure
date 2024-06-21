package datasource

import (
	"errors"
	"template-ulamm-backend-go/utils"

	"gorm.io/gorm"
)

type Datasource struct {
	GormDB *gorm.DB

	isUseSqlServerPrimaryDb bool
}

var dataSource *Datasource = &Datasource{}

func (ds *Datasource) PingDB() error {
	if pinger, ok := ds.GormDB.ConnPool.(interface{ Ping() error }); ok {
		if err := pinger.Ping(); err != nil {
			return err
		}
	} else {
		return errors.New("failed to ping")
	}

	return nil
}

func NewDatasource() error {
	err := NewGormDb(utils.GetConfig().SqlServer.Hosts[0])
	if err != nil {
		return err
	}

	return nil
}

func GetDatasource() *Datasource {
	return dataSource
}
