package entity

import "time"

type Program struct {
	IDProgram    int       `gorm:"primaryKey;column:id_program"`
	ProgramName  string    `gorm:"column:program_name"`
	IsActive     int       `gorm:"column:is_active"`
	CreatedBy    string    `gorm:"column:created_by"`
	CreatedTime  time.Time `gorm:"column:created_time;autoCreateTime"`
	ModifiedBy   string    `gorm:"column:modified_by"`
	ModifiedTime time.Time `gorm:"column:modified_time;autoUpdateTime"`
}

func (Program) TableName() string {
	return "Mst_Program"
}
