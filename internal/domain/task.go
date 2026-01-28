package domain

import "time"

type Task struct {
	Header      string
	Description string
	IsCompleted bool

	CreatedAt   time.Time
	CompletedAt *time.Time
}
