package usecase

import (
	"time"

	"github.com/yakupovdev/ToDoList/model"
	"github.com/yakupovdev/ToDoList/repository"
)

type TaskUsecase struct {
	repo *repository.TaskRepository
}

func NewTaskUsecase(repo *repository.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		repo: repo,
	}
}

func (uc *TaskUsecase) AddTask(header, description string) error {
	task := model.Task{
		Header:       header,
		Description:  description,
		IsCompleted:  false,
		CreationTime: time.Now(),
		CompleteTime: time.Time{},
	}

	uc.repo.AddTask(task)

	return nil
}

func (uc *TaskUsecase) GetTasks() ([]model.Task, error) {
	return uc.repo.GetTasks(), nil
}

func (uc *TaskUsecase) GetUncompletedTasks() []model.Task {
	return uc.repo.GetUncompletedTasks()
}

func (uc *TaskUsecase) GetTask(header string) (model.Task, error) {
	return uc.repo.GetTask(header)
}

func (uc *TaskUsecase) RemoveTask(header string) error {
	return uc.repo.RemoveTask(header)
}
