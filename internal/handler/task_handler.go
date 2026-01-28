package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	model2 "github.com/yakupovdev/ToDoList/internal/model"
	"github.com/yakupovdev/ToDoList/internal/repository"
	"github.com/yakupovdev/ToDoList/internal/usecase"
)

type TaskHandler struct {
	uc *usecase.TaskUsecase
}

func NewTaskHandler(uc *usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{
		uc: uc,
	}
}

func (th *TaskHandler) HandleAddTask(w http.ResponseWriter, r *http.Request) {
	var req model2.TaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errResponse := model2.ErrorResponse{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errResponse.ToString(), http.StatusBadRequest)
		return
	}

	err := req.Validate()
	if err != nil {
		errResponse := model2.ErrorResponse{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errResponse.ToString(), http.StatusBadRequest)
		return
	}

	task, err := th.uc.AddTask(req.Header, req.Description)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrTaskAlreadyExists):
			errResponse := model2.ErrorResponse{
				Message: err.Error(),
				Time:    time.Now(),
			}
			http.Error(w, errResponse.ToString(), http.StatusConflict)
		default:
			errResponse := model2.ErrorResponse{
				Message: err.Error(),
				Time:    time.Now(),
			}
			http.Error(w, errResponse.ToString(), http.StatusInternalServerError)
		}
		return
	}

	res := model2.TaskResponse{
		Header:      task.Header,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		CompletedAt: task.CompletedAt,
		CreatedAt:   task.CreatedAt,
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(res.ToByteArray())
	if err != nil {
		log.Println("failed to write http response", err)
		return
	}
}

func (th *TaskHandler) HandleChangeCompleteStatusTask(w http.ResponseWriter, r *http.Request) {
	header := mux.Vars(r)["header"]
	var req model2.CompleteTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errResponse := model2.ErrorResponse{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errResponse.ToString(), http.StatusBadRequest)
		return
	}
	task, err := th.uc.ChangeCompleteStatusTask(header, req.IsCompleted)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrTaskNotFound):
			errResponse := model2.ErrorResponse{
				Message: err.Error(),
				Time:    time.Now(),
			}
			http.Error(w, errResponse.ToString(), http.StatusNotFound)
		default:
			errResponse := model2.ErrorResponse{
				Message: err.Error(),
				Time:    time.Now(),
			}
			http.Error(w, errResponse.ToString(), http.StatusInternalServerError)
		}
		return
	}
	res := model2.TaskResponse{
		Header:      task.Header,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		CompletedAt: task.CompletedAt,
		CreatedAt:   task.CreatedAt,
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res.ToByteArray())
	if err != nil {
		log.Println("failed to write http response", err)
		return
	}
}

func (th *TaskHandler) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := th.uc.GetTasks()
	b, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		log.Println("failed to marshal json", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		log.Println("failed to write http response", err)
		return
	}
}

func (th *TaskHandler) HandleGetUncompletedTasks(w http.ResponseWriter, r *http.Request) {
	tasks := th.uc.GetUncompletedTasks()
	b, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		log.Println("failed to marshal json", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		log.Println("failed to write http response", err)
		return
	}
}

func (th *TaskHandler) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	header := mux.Vars(r)["header"]
	task, err := th.uc.GetTask(header)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrTaskNotFound):
			errResponse := model2.ErrorResponse{
				Message: err.Error(),
				Time:    time.Now(),
			}
			http.Error(w, errResponse.ToString(), http.StatusNotFound)
		default:
			errResponse := model2.ErrorResponse{
				Message: err.Error(),
				Time:    time.Now(),
			}
			http.Error(w, errResponse.ToString(), http.StatusInternalServerError)
		}
		return
	}

	res := model2.TaskResponse{
		Header:      task.Header,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		CompletedAt: task.CompletedAt,
		CreatedAt:   task.CreatedAt,
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res.ToByteArray())
	if err != nil {
		log.Println("failed to write http response", err)
		return
	}

}

func (th *TaskHandler) HandleRemoveTask(w http.ResponseWriter, r *http.Request) {
	header := mux.Vars(r)["header"]
	err := th.uc.RemoveTask(header)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrTaskNotFound):
			errResponse := model2.ErrorResponse{
				Message: err.Error(),
				Time:    time.Now(),
			}
			http.Error(w, errResponse.ToString(), http.StatusNotFound)
		default:
			errResponse := model2.ErrorResponse{
				Message: err.Error(),
				Time:    time.Now(),
			}
			http.Error(w, errResponse.ToString(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("{}"))
}
