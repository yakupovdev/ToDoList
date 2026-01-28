package repository

import (
	"sync"
	"time"

	"github.com/yakupovdev/ToDoList/internal/model"
	"github.com/yakupovdev/ToDoList/internal/storage"
)

type Repository interface {
	AddTask(model.Task) (model.Task, error)
	GetTask(string) (model.Task, error)
	RemoveTask(string) error
	GetTasks() []model.Task
	GetUncompletedTasks() []model.Task
	ChangeCompleteStatusTask(string, bool) (model.Task, error)
}

type TaskRepository struct {
	mtx     sync.RWMutex
	storage *storage.Storage
}

func NewRepositoryStorage(storage *storage.Storage) *TaskRepository {
	return &TaskRepository{
		storage: storage,
		mtx:     sync.RWMutex{},
	}
}

func (s *TaskRepository) AddTask(task model.Task) (model.Task, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.storage.Tasks[task.Header]; ok {
		return model.Task{}, ErrTaskAlreadyExists
	}
	s.storage.Tasks[task.Header] = task
	return task, nil
}

func (s *TaskRepository) GetTask(header string) (model.Task, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	task, ok := s.storage.Tasks[header]
	if !ok {
		return model.Task{}, ErrTaskNotFound
	}
	return task, nil
}

func (s *TaskRepository) GetTasks() []model.Task {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	tasks := make([]model.Task, 0, len(s.storage.Tasks))
	for _, task := range s.storage.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (s *TaskRepository) GetUncompletedTasks() []model.Task {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	tasks := make([]model.Task, 0, len(s.storage.Tasks))
	for _, task := range s.storage.Tasks {
		if !task.IsCompleted {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func (s *TaskRepository) ChangeCompleteStatusTask(header string, status bool) (model.Task, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	task, ok := s.storage.Tasks[header]
	if !ok {
		return model.Task{}, ErrTaskNotFound
	}
	task.IsCompleted = status
	if status {
		task.CompletedAt = time.Now()
	} else {
		task.CompletedAt = time.Time{}
	}
	s.storage.Tasks[header] = task
	return task, nil
}

func (s *TaskRepository) RemoveTask(header string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	_, ok := s.storage.Tasks[header]
	if !ok {
		return ErrTaskNotFound
	}

	delete(s.storage.Tasks, header)
	return nil
}
