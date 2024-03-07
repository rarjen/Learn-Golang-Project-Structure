package repository

import (
	"context"
	"errors"
	"fmt"
	"template-ulamm-backend-go/domain"
	"template-ulamm-backend-go/domain/entity"
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/utils"
	"template-ulamm-backend-go/utils/constantvar"
	"time"

	"gorm.io/plugin/dbresolver"
)

type SyncDataRepository interface {
	GetDummyData() error
	SyncDataGetConnectionString(
		kodeUnit string,
	) (*entity.DBUnitConnectionEntity, error)
	ScoringUsers(
		ctx context.Context,
	) ([]entity.SyncDataScoringParameterAPIEntity, error)
	CheckDebiturIDUnit(
		ctx context.Context,
		stringConnection string,
		nasabahData domain.CheckDebiturIDRequest) ([]entity.CheckDebiturIDEntity, error)
}

type syncDataRepository struct {
	DB                       *datasource.Datasource
	ExternalSourceRepository ExternalSourceRepository
}

func NewSyncDataRepository(
	ds *datasource.Datasource,
	externalSourceRepo ExternalSourceRepository,
) SyncDataRepository {
	return &syncDataRepository{
		DB:                       ds,
		ExternalSourceRepository: externalSourceRepo,
	}
}

func (sDR *syncDataRepository) SyncDataGetConnectionString(
	kodeUnit string,
) (*entity.DBUnitConnectionEntity, error) {
	result := &entity.DBUnitConnectionEntity{}
	if err := sDR.DB.Db.Raw(
		fmt.Sprintf(
			`
		SELECT
			RTRIM(KODE_CAB) KODE_CAB,
			RTRIM(INISIAL) INISIAL,
			RTRIM (NAMA_CAB) NAMA_CAB,
			RTRIM(DATABASE_CAB) DATABASE_CAB,
			RTRIM(IP_CAB) IP_CAB,
			RTRIM(INISIAL_CAB_INDUK) INISIAL_CAB_INDUK,
			RTRIM (NAMA_CAB_INDUK) NAMA_CAB_INDUK,
			RTRIM(USER_CAB) USER_CAB,
			RTRIM(PWD_CAB) PWD_CAB,
			RTRIM(mw.KodeWilayah) KODE_WILAYAH
		FROM
			kodecabang
		left join master_wilayah mw on
			mw.KodeCabang = kodecabang.INISIAL_CAB_INDUK
		WHERE
			tipe_cab = 2
			AND kode_cab = '%s'
		`,
			kodeUnit,
		),
	).Scan(&result).Error; err != nil {
		return nil, err
	}
	if *result == (entity.DBUnitConnectionEntity{}) {
		return nil, errors.New(constantvar.HTTP_RESPONSE_DATA_NOT_FOUND)
	}
	return result, nil
}

func (sDR *syncDataRepository) GetDummyData() error {
	return nil
}

func (sDR *syncDataRepository) ScoringUsers(
	ctx context.Context,
) ([]entity.SyncDataScoringParameterAPIEntity, error) {
	userEntities := make([]entity.SyncDataScoringParameterAPIEntity, 0)
	if err := sDR.DB.Db.
		WithContext(ctx).
		Clauses(dbresolver.Use(constantvar.SECONDARY_DB_MIS)).
		Table("scoring_parameter_API").
		Select(`TOP 10 *`).
		Scan(&userEntities).Error; err != nil {
		return nil, err
	}
	return userEntities, nil
}

func (sDR *syncDataRepository) CheckDebiturIDUnit(
	ctx context.Context,
	stringConnection string,
	nasabahData domain.CheckDebiturIDRequest,
) ([]entity.CheckDebiturIDEntity, error) {

	result := make([]entity.CheckDebiturIDEntity, 0)
	sqlQuery := fmt.Sprintf(
		`
			SELECT
				RTRIM(nasabah_id) nasabah_id,
				tgllahir ,
				TGL_AKTE_AKHIR,
				RTRIM(no_id) no_id
			FROM
				nasabah
			WHERE
				nasabah_id = '%s';
		`,
		nasabahData.NasabahID,
	)

	dbResult, err := sDR.ExternalSourceRepository.QueryDB(
		ctx,
		stringConnection,
		sqlQuery,
	)

	if err != nil {
		return nil, err
	}

	if len(dbResult) == 0 {
		return nil, errors.New(constantvar.HTTP_RESPONSE_DATA_NOT_FOUND)
	}

	for _, v := range dbResult {
		tmpData := entity.CheckDebiturIDEntity{}
		if _, ok := v["nasabah_id"]; ok {
			tmpData.DebiturID = v["nasabah_id"].(string)
		}

		if value, ok := v["TGL_AKTE_AKHIR"]; ok {
			if value != nil {
				tmpNullableBirthCertDate := v["TGL_AKTE_AKHIR"].(time.Time)
				birthDateCert := utils.TimeToSQLDateConverter(tmpNullableBirthCertDate)
				tmpData.BirthCertDate = &birthDateCert
			}
		}

		if value, ok := v["tgllahir"]; ok {
			if value == nil {
				tmpNullableTglLahir := v["tgllahir"].(time.Time)
				tmpStringTglLahir := utils.TimeToSQLDateConverter(tmpNullableTglLahir)
				tmpData.BirthDate = &tmpStringTglLahir
			}
		}

		if _, ok := v["no_id"]; ok {
			tmpData.NumberID = v["no_id"].(string)
		}

		result = append(result, tmpData)
	}

	return result, nil
}
