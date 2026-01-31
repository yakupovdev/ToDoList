package repository

import (
	"sync"
	"time"

	"github.com/yakupovdev/ToDoList/internal/domain"
)

type TaskRepository struct {
	tasks map[string]domain.Task
	mtx   sync.RWMutex
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[string]domain.Task),
		mtx:   sync.RWMutex{},
	}
}

func (s *TaskRepository) AddTask(task domain.Task) (domain.Task, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.tasks[task.Header]; ok {
		return domain.Task{}, domain.ErrTaskAlreadyExists
	}
	s.tasks[task.Header] = task
	return task, nil
}

func (s *TaskRepository) GetTask(header string) (domain.Task, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	task, ok := s.tasks[header]
	if !ok {
		return domain.Task{}, domain.ErrTaskNotFound
	}
	return task, nil
}

func (s *TaskRepository) GetTasks() []domain.Task {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	tasks := make([]domain.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (s *TaskRepository) GetUncompletedTasks() []domain.Task {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	tasks := make([]domain.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		if !task.IsCompleted {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func (s *TaskRepository) ChangeCompleteStatusTask(header string, status bool) (domain.Task, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	task, ok := s.tasks[header]
	if !ok {
		return domain.Task{}, domain.ErrTaskNotFound
	}
	task.IsCompleted = status
	if status {
		now := time.Now()
		task.CompletedAt = &now
	} else {
		task.CompletedAt = nil
	}
	s.tasks[header] = task
	return task, nil
}

func (s *TaskRepository) RemoveTask(header string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	_, ok := s.tasks[header]
	if !ok {
		return domain.ErrTaskNotFound
	}

	delete(s.tasks, header)
	return nil
}
