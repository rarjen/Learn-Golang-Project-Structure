package domain

// CheckDebiturID related model

type CheckDebiturIDRequest struct {
	NasabahID string `json:"nasabahID" binding:"required" example:"HC.000251"`
	KodeUnit  string `json:"unitCode" binding:"required" example:"HC"`
}

// GetUnitBranchCode related model
type GetUnitBranchCodeRequest struct {
	Limit     int64  `json:"limit" binding:"required" example:"10"`
	Condition string `json:"condition" example:"kode_cab like '%Y%'"`
}

// Request model when do HTTP request to get decrypted password
type DecodeDBPasswordRequest struct {
	Unit     string `json:"unit"`
	Password string `json:"password"`
}