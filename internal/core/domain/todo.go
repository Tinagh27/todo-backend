package domain

import "time"

type Todo struct {
	ID          int64
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
