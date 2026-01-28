package model

import (
	"encoding/json"
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

	CompletedAt time.Time `json:"completed_at"`
	CreatedAt   time.Time `json:"created_at"`
}

func (addTaskResponse *TaskResponse) ToByteArray() []byte {
	value, err := json.MarshalIndent(addTaskResponse, "", "	")
	if err != nil {
		panic(err)
	}
	return value
}

type CompleteTaskRequest struct {
	IsCompleted bool `json:"is_completed"`
}
