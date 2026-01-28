package http

import (
	"errors"
	"time"
)

type TaskRequest struct {
	Header      string `json:"header"`
	Description string `json:"description"`
}

func (addTaskRequest *TaskRequest) Validate() error {
	if addTaskRequest.Header == "" {
		return errors.New("header is required")
	}

	if addTaskRequest.Description == "" {
		return errors.New("description is required")
	}

	return nil
}

type TaskResponse struct {
	Header      string `json:"header"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`

	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type CompleteTaskRequest struct {
	IsCompleted bool `json:"is_completed"`
}

type ErrorResponse struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
