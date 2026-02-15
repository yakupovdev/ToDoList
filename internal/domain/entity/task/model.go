package task

import (
	"strings"
	"time"
)

type Task struct {
	Header      string
	Description string
	IsCompleted bool

	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(header, description string) (*Task, error) {
	header = strings.TrimSpace(header)
	description = strings.TrimSpace(description)

	if header == "" {
		return nil, ErrHeaderRequired
	}

	if description == "" {
		return nil, ErrDescriptionRequired
	}

	return &Task{
		Header:      header,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}, nil
}

func RestoreTask(header, description string, isCompleted bool, createdAt time.Time, completedAt *time.Time) *Task {
	return &Task{
		Header:      header,
		Description: description,
		IsCompleted: isCompleted,
		CreatedAt:   createdAt,
		CompletedAt: completedAt,
	}
}
