package datasource

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDb(host string) error {
	sqlDialector := sqlserver.New(sqlserver.Config{
		DSN: host,
	})

	gormdb, err := gorm.Open(sqlDialector, &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.LogLevel(4)),
	})
	if err != nil {
		return err
	}

	dataSource.GormDB = gormdb

	return nil
}
