package datasource

import (
	"template-ulamm-backend-go/utils/config"
	"template-ulamm-backend-go/utils/constantvar"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type Datasource struct {
	Db *gorm.DB
}

func NewGORMDB(conf config.Config) (*gorm.DB, error) {
	connMMSKonve, err := mmsKonveDB(conf)
	if err != nil {
		return nil, err
	}
	sqlConveDialector := sqlserver.New(sqlserver.Config{
		Conn: connMMSKonve,
	})

	connMMSMIS, err := mmsMISDB(conf)
	if err != nil {
		return nil, err
	}
	sqlMISDialector := sqlserver.New(sqlserver.Config{
		Conn: connMMSMIS,
	})

	db, err := gorm.Open(sqlConveDialector, &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.LogLevel(4)),
	})
	if err != nil {
		return nil, err
	}
	db.Use(
		dbresolver.
			Register(dbresolver.Config{
				Sources: []gorm.Dialector{
					sqlMISDialector,
				},
			}, constantvar.SECONDARY_DB_MIS,
			),
	)
	return db, nil
}
