package entity

/*
	entity format name
	BusinessDomainTableEntity
	example:
	the business domain is syncData
	the table is ScoringParameterAPi
	so the struct name is SyncDataScoringParameterAPI

	if the query is complicated
	BusinessDomainView
	example:
	the business domain is syncData
	the query is to get Scoring for each user
	so the struct name is SyncDataScoringUserEntity
*/

type SyncDataScoringParameterAPIEntity struct {
	ID            int64    `gorm:"column:ID"`
	IDName        *string  `gorm:"column:ID_Nama"`
	IDOrder       *int64   `gorm:"column:ID_Order"`
	IDMapping     *string  `gorm:"column:ID_Mapping"`
	IDField       *string  `gorm:"column:ID_field"`
	Description   *string  `gorm:"column:Deskripsi"`
	Weight        *int64   `gorm:"column:Bobot"`
	BottomLimit   *float64 `gorm:"column:Batas_Bawah"`
	UpperLimit    *float64 `gorm:"column:Batas_Atas"`
	Category      *string  `gorm:"column:Kategori"`
	ParameterType *int64   `gorm:"column:Jenis_Parameter"`
}

type CheckDebiturIDEntity struct {
	DebiturID     string  `gorm:"column:nasabah_id"`
	BirthDate     *string `gorm:"column:tgllahir"`
	BirthCertDate *string `gorm:"column:TGL_AKTE_AKHIR"`
	NumberID      string  `gorm:"column:no_id"`
}

// Get DB Unit connection
type DBUnitConnectionEntity struct {
	BranchCode           string `gorm:"column:KODE_CAB"`
	Initials             string `gorm:"column:INISIAL"`
	BranchName           string `gorm:"column:NAMA_CAB"`
	BranchDatabase       string `gorm:"column:DATABASE_CAB"`
	BranchIP             string `gorm:"column:IP_CAB"`
	ParentBranchInitials string `gorm:"column:INISIAL_CAB_INDUK"`
	ParentBranchName     string `gorm:"column:NAMA_CAB_INDUK"`
	BranchUser           string `gorm:"column:USER_CAB"`
	BranchPassword       string `gorm:"column:PWD_CAB"`
	RegionCode           string `gorm:"column:KODE_WILAYAH"`
}
