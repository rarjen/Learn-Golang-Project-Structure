package entity

import "time"

type Ping struct {
	CurrentDate time.Time `db:"current_date"`
}
