package response

import "time"

type PingResponse struct {
	Status      string    `json:"status"`
	Message     any       `json:"message"`
	CurrentDate time.Time `json:"current_date"`
}

type Pagination struct {
	Count int64 `json:"count"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
}

type PaginatedData[T any] struct {
	Data       []T `json:"data"`
	Pagination `json:"pagination"`
}
