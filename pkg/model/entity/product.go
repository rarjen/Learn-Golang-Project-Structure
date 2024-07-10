package entity

import "time"

type Product struct {
	IDProduct          int `gorm:"primaryKey"`
	ProductName        string
	ProductCode        string
	InterestRate       float64
	InterestRateAnnual float64
	LimitLoanLower     float64
	LimitLoanUpper     float64
	TimePeriodLower    int
	TimePeriodUpper    int
	IsActive           int
	CreatedBy          string
	CreatedTime        time.Time
	ModifiedBy         string
	ModifiedTime       time.Time
}

func (Product) TableName() string {
	return "Mst_Product"
}
