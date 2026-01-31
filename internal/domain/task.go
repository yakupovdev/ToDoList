package domain

import "time"

type Task struct {
	Header      string
	Description string
	IsCompleted bool

	CreatedAt   time.Time
	CompletedAt *time.Time
}

type TaskRepository interface {
	AddTask(Task) (Task, error)
	GetTask(string) (Task, error)
	RemoveTask(string) error
	GetTasks() []Task
	GetUncompletedTasks() []Task
	ChangeCompleteStatusTask(string, bool) (Task, error)
}
