package repository

import (
	"sync"
	"time"

	"github.com/yakupovdev/ToDoList/internal/domain/entity/task"
)

type TaskRepository struct {
	tasks map[string]task.Task
	mtx   sync.RWMutex
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[string]task.Task),
		mtx:   sync.RWMutex{},
	}
}

func (s *TaskRepository) AddTask(t task.Task) (task.Task, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.tasks[t.Header]; ok {
		return task.Task{}, task.ErrTaskAlreadyExists
	}
	s.tasks[t.Header] = t
	return t, nil
}

func (s *TaskRepository) GetTask(header string) (task.Task, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	t, ok := s.tasks[header]
	if !ok {
		return task.Task{}, task.ErrTaskNotFound
	}
	return t, nil
}

func (s *TaskRepository) GetTasks() []task.Task {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	tasks := make([]task.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (s *TaskRepository) GetUncompletedTasks() []task.Task {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	tasks := make([]task.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		if !task.IsCompleted {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func (s *TaskRepository) ChangeCompleteStatusTask(header string, status bool) (task.Task, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	t, ok := s.tasks[header]
	if !ok {
		return task.Task{}, task.ErrTaskNotFound
	}
	t.IsCompleted = status
	if status {
		now := time.Now()
		t.CompletedAt = &now
	} else {
		t.CompletedAt = nil
	}
	s.tasks[header] = t
	return t, nil
}

func (s *TaskRepository) RemoveTask(header string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	_, ok := s.tasks[header]
	if !ok {
		return task.ErrTaskNotFound
	}

	delete(s.tasks, header)
	return nil
}
