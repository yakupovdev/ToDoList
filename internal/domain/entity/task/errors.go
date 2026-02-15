package task

import "errors"

var (
	ErrTaskNotFound        = errors.New("task not found")
	ErrTaskAlreadyExists   = errors.New("task with this header already exists")
	ErrHeaderRequired      = errors.New("header is required")
	ErrDescriptionRequired = errors.New("description is required")
)
