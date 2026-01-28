package usecase

import (
	"time"

	"github.com/yakupovdev/ToDoList/internal/model"
	"github.com/yakupovdev/ToDoList/internal/repository"
)

type TaskUsecase struct {
	repo *repository.TaskRepository
}

func NewTaskUsecase(repo *repository.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		repo: repo,
	}
}

func (uc *TaskUsecase) AddTask(header, description string) (model.Task, error) {
	task := model.Task{
		Header:      header,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	if _, err := uc.repo.AddTask(task); err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (uc *TaskUsecase) GetTasks() []model.Task {
	return uc.repo.GetTasks()
}

func (uc *TaskUsecase) GetUncompletedTasks() []model.Task {
	return uc.repo.GetUncompletedTasks()
}

func (uc *TaskUsecase) GetTask(header string) (model.Task, error) {
	return uc.repo.GetTask(header)
}

func (uc *TaskUsecase) ChangeCompleteStatusTask(header string, isCompleted bool) (model.Task, error) {
	task, err := uc.repo.ChangeCompleteStatusTask(header, isCompleted)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (uc *TaskUsecase) RemoveTask(header string) error {
	return uc.repo.RemoveTask(header)
}
