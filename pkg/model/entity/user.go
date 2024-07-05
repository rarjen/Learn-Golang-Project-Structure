package entity

import "time"

type User struct {
	ID           string
	IDEmployee   string
	Username     string
	Name         string
	IsActive     int
	CreatedBy    string
	CreatedTime  time.Time
	ModifiedBy   string
	ModifiedTime time.Time
}

func (User) TableName() string {
	return "Mst_User"
}
