package domain

// CheckDebiturID related model

type CheckDebiturIDRequest struct {
	NasabahID string `json:"nasabahID" binding:"required"`
	KodeUnit  string `json:"unitCode" binding:"required"`
}

// Request model when do HTTP request to get decrypted password
type DecodeDBPasswordRequest struct {
	Unit     string `json:"unit"`
	Password string `json:"password"`
}