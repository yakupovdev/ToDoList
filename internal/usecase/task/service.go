package task

import (
	"time"

	"github.com/yakupovdev/ToDoList/internal/domain/entity/task"
)

type TaskUsecase struct {
	userRepo Repository
}

func NewTaskUsecase(repo Repository) *TaskUsecase {
	return &TaskUsecase{
		userRepo: repo,
	}
}

func (uc *TaskUsecase) AddTask(header, description string) (task.Task, error) {
	t := task.Task{
		Header:      header,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	if _, err := uc.userRepo.AddTask(t); err != nil {
		return task.Task{}, err
	}

	return t, nil
}

func (uc *TaskUsecase) GetTasks() []task.Task {
	return uc.userRepo.GetTasks()
}

func (uc *TaskUsecase) GetUncompletedTasks() []task.Task {
	return uc.userRepo.GetUncompletedTasks()
}

func (uc *TaskUsecase) GetTask(header string) (task.Task, error) {
	return uc.userRepo.GetTask(header)
}

func (uc *TaskUsecase) ChangeCompleteStatusTask(header string, isCompleted bool) (task.Task, error) {
	t, err := uc.userRepo.ChangeCompleteStatusTask(header, isCompleted)
	if err != nil {
		return task.Task{}, err
	}
	return t, nil
}

func (uc *TaskUsecase) RemoveTask(header string) error {
	return uc.userRepo.RemoveTask(header)
}
