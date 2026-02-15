package task

import "github.com/yakupovdev/ToDoList/internal/domain/entity/task"

type Repository interface {
	AddTask(task task.Task) (task.Task, error)
	GetTask(string) (task.Task, error)
	RemoveTask(string) error
	GetTasks() []task.Task
	GetUncompletedTasks() []task.Task
	ChangeCompleteStatusTask(string, bool) (task.Task, error)
}
