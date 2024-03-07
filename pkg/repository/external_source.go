package repository

import (
	"context"
	"database/sql"
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/utils/config"
	"template-ulamm-backend-go/utils/constantvar"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*
	Will be used when invoke external call (API/DB)
	Don't use this repository multiple times in 1 or more Use Case
	Kindly to create a new repository for each usecase
	1 use case using 1 or more repo ✅
	1 repo used by more than 1 use case ❎
	1 utility repo used by 1 or more repo (similar with utils) ✅
*/

type ExternalSourceRepository interface {
	QueryDB(
		ctx context.Context,
		stringConnection,
		query string,
	) ([]map[string]interface{}, error)
}

type externalSourceRepository struct {
	DB *datasource.Datasource
}

func NewExternalSourceRepository(ds *datasource.Datasource) ExternalSourceRepository {
	return &externalSourceRepository{
		DB: ds,
	}
}

func (gR *externalSourceRepository) newGormDB(
	dbCon string,
) (*gorm.DB, error) {
	connMMSKonveDB, err := gR.openDBCon(dbCon)
	if err != nil {
		return nil, err
	}
	sqlDialector := sqlserver.New(sqlserver.Config{
		Conn: connMMSKonveDB,
	})

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.LogLevel(constantvar.DB_LOG_LEVEL_PRODUCTION)),
	}

	if config.C.STAGE == constantvar.STAGE_DEVELOPMENT {
		gormConfig.Logger = logger.Default.LogMode(logger.LogLevel(constantvar.DB_LOG_LEVEL_DEVELOPMENT))
	}

	db, err := gorm.Open(sqlDialector, gormConfig)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (gR *externalSourceRepository) openDBCon(
	dbCon string,
) (*sql.DB, error) {
	// dsn := dbCon
	sqlDB, err := sql.Open("sqlserver", dbCon)
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(2 * time.Minute)
	sqlDB.SetMaxOpenConns(2)
	sqlDB.SetMaxIdleConns(2)
	return sqlDB, nil
}

func (eSR *externalSourceRepository) QueryDB(
	ctx context.Context,
	stringConnection,
	query string,
) ([]map[string]interface{}, error) {
	dB, err := eSR.newGormDB(stringConnection)
	if err != nil {
		return nil, err
	}

	var resultDB []map[string]interface{}
	
	if err := dB.WithContext(ctx).Raw(query).Scan(&resultDB).Error; err != nil {
		return nil, err
	}

	return resultDB, nil
}
