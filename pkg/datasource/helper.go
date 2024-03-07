package datasource

import (
	"database/sql"
	"template-ulamm-backend-go/utils/config"
	"time"
)

/* connection to multiple db will created here */

func mmsKonveDB(conf config.Config) (*sql.DB, error) {
	sqlDB, err := sql.Open("sqlserver", conf.DBDSNKonve)
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	return sqlDB, nil
}

func mmsMISDB(conf config.Config) (*sql.DB, error) {
	sqlDB, err := sql.Open("sqlserver", conf.DBDSNMIS)
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	return sqlDB, nil
}
