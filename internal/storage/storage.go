package storage

import (
	"github.com/yakupovdev/ToDoList/internal/model"
)

type Storage struct {
	Tasks map[string]model.Task
}

func NewStorage() *Storage {
	return &Storage{
		Tasks: make(map[string]model.Task),
	}
}
