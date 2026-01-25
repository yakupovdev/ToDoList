package repository

import (
	"sync"
	"time"

	"github.com/yakupovdev/ToDoList/model"
	"github.com/yakupovdev/ToDoList/storage"
)

type Repository interface {
	AddTask(task model.Task)
	GetTask(header string) (model.Task, error)
	RemoveTask(header string) error
	GetTasks() []model.Task
	GetUncompletedTasks() []model.Task
	CompleteTask(header string) error
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

func (s *TaskRepository) AddTask(task model.Task) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.storage.Tasks[task.Header] = task
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

func (s *TaskRepository) CompleteTask(header string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	task, ok := s.storage.Tasks[header]
	if !ok {
		return ErrTaskNotFound
	}
	task.IsCompleted = true
	task.CompleteTime = time.Now()

	s.storage.Tasks[header] = task
	return nil
}

func (s *TaskRepository) RemoveTask(header string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	_, err := s.GetTask(header)
	if err != nil {
		return ErrTaskNotFound
	}

	delete(s.storage.Tasks, header)
	return nil
}
