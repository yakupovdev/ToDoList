package usecase

import (
	"time"

	"github.com/yakupovdev/ToDoList/internal/domain"
)

type TaskUsecase struct {
	repo domain.TaskRepository
}

func NewTaskUsecase(repo domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		repo: repo,
	}
}

func (uc *TaskUsecase) AddTask(header, description string) (domain.Task, error) {
	task := domain.Task{
		Header:      header,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	if _, err := uc.repo.AddTask(task); err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (uc *TaskUsecase) GetTasks() []domain.Task {
	return uc.repo.GetTasks()
}

func (uc *TaskUsecase) GetUncompletedTasks() []domain.Task {
	return uc.repo.GetUncompletedTasks()
}

func (uc *TaskUsecase) GetTask(header string) (domain.Task, error) {
	return uc.repo.GetTask(header)
}

func (uc *TaskUsecase) ChangeCompleteStatusTask(header string, isCompleted bool) (domain.Task, error) {
	task, err := uc.repo.ChangeCompleteStatusTask(header, isCompleted)
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func (uc *TaskUsecase) RemoveTask(header string) error {
	return uc.repo.RemoveTask(header)
}
