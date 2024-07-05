package entity

import "time"

type City struct {
	ID         string `gorm:"primaryKey"`
	IDCity     int
	IDProv     int
	CityName   string
	IsActive   int
	CreatedBy  string
	CreatedAt  time.Time
	ModifiedBy string
	ModifiedAt time.Time
}

func (City) TableName() string {
	return "Mst_City"
}
