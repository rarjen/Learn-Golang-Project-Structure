package domain

/*
	response format name
	BusinessDomainTableresponse
	example:
	the business domain is syncData
	the table is ScoringParameterAPi
	so the struct name is SyncDataScoringParameterAPIResponse

	if the query is complicated
	BusinessDomainView
	example:
	the business domain is syncData
	the query is to get Scoring for each user
	so the struct name is SyncDataScoringUserResponse

*/

// ScoringUsers controller related

type SyncDataScoringParameterAPIResponse struct {
	Status      int                               `json:"status" example:"100"`
	Description string                            `json:"description" example:"success"`
	Data        []SyncDataScoringParameterAPIData `json:"data"`
}

type SyncDataScoringParameterAPIData struct {
	ID          int64    `json:"id" example:""`
	BottomLimit *float64 `json:"batas_bawah" example:""`
	UpperLimit  *float64 `json:"batas_atas" example:""`
	Category    *string  `json:"kategory" example:""`
}

// CheckDebiturID related response
type CheckDebiturIDResponse struct {
	Status      int                  `json:"status" example:"100"`
	Description string               `json:"description" example:"success"`
	Data        []CheckDebiturIDData `json:"data"`
}

type CheckDebiturIDData struct {
	NumberID  string  `json:"numberID"`
	BirthDate *string `json:"birthDate"`
}

// GetUnitBranchCode (KodeCabangUnit)

type GetUnitBranchCodeResponse struct {
	Status      int                     `json:"status" example:"100"`
	Description string                  `json:"description" example:"success"`
	Data        []GetUnitBranchCodeData `json:"data"`
}

type GetUnitBranchCodeData struct {
	BranchCode   string `json:"kodeCabang" example:"YR"`
	BranchName   string `json:"namaCabang" example:"Unit Slipi"`
	BranchIP     string `json:"ipCabang" example:"10.61.4.102"`
	BranchDBName string `json:"namaDBCabang" example:"MMS-JKT-YR"`
}

/*
	utility response
*/

// DecodeDBPassword to external API
type DecodeDBPasswordResponse struct {
	DecryptedString *string `json:"decryptedString"`
}
