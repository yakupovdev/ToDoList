package model

import "time"

type Task struct {
	Header      string
	Description string
	IsCompleted bool

	CreationTime time.Time
	CompleteTime time.Time
}
