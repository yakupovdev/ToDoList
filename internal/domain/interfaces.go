package domain

type TaskRepository interface {
	AddTask(Task) (Task, error)
	GetTask(string) (Task, error)
	RemoveTask(string) error
	GetTasks() []Task
	GetUncompletedTasks() []Task
	ChangeCompleteStatusTask(string, bool) (Task, error)
}
