package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	http2 "github.com/yakupovdev/ToDoList/internal/delivery/http/dto"
	"github.com/yakupovdev/ToDoList/internal/domain"
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
	var req http2.TaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errResponse := http2.ErrorResponse{
			Message: err.Error(),
			Time:    time.Now(),
		}

		th.respondWithJSON(w, http.StatusBadRequest, errResponse)
		return
	}

	err := req.Validate()
	if err != nil {
		errResponse := http2.ErrorResponse{
			Message: err.Error(),
			Time:    time.Now(),
		}

		th.respondWithJSON(w, http.StatusBadRequest, errResponse)
		return
	}

	task, err := th.uc.AddTask(req.Header, req.Description)
	if err != nil {
		errResponse := http2.ErrorResponse{
			Message: err.Error(),
			Time:    time.Now(),
		}
		switch {
		case errors.Is(err, domain.ErrTaskAlreadyExists):
			th.respondWithJSON(w, http.StatusConflict, errResponse)
		default:
			th.respondWithJSON(w, http.StatusInternalServerError, errResponse)
		}
		return
	}

	res := http2.TaskResponse{
		Header:      task.Header,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		CompletedAt: task.CompletedAt,
		CreatedAt:   task.CreatedAt,
	}

	th.respondWithJSON(w, http.StatusCreated, res)
}

func (th *TaskHandler) HandleChangeCompleteStatusTask(w http.ResponseWriter, r *http.Request) {
	header := mux.Vars(r)["header"]
	var req http2.CompleteTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errResponse := http2.ErrorResponse{
			Message: err.Error(),
			Time:    time.Now(),
		}
		log.Println(err)
		th.respondWithJSON(w, http.StatusBadRequest, errResponse)
		return
	}
	task, err := th.uc.ChangeCompleteStatusTask(header, req.IsCompleted)
	if err != nil {
		errResponse := http2.ErrorResponse{
			Message: err.Error(),
			Time:    time.Now(),
		}
		switch {
		case errors.Is(err, domain.ErrTaskNotFound):
			th.respondWithJSON(w, http.StatusNotFound, errResponse)
		default:
			th.respondWithJSON(w, http.StatusInternalServerError, errResponse)
		}
		return
	}
	res := http2.TaskResponse{
		Header:      task.Header,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		CompletedAt: task.CompletedAt,
		CreatedAt:   task.CreatedAt,
	}

	th.respondWithJSON(w, http.StatusOK, res)
}

func (th *TaskHandler) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := th.uc.GetTasks()

	th.respondWithJSON(w, http.StatusOK, tasks)
}

func (th *TaskHandler) HandleGetUncompletedTasks(w http.ResponseWriter, r *http.Request) {
	tasks := th.uc.GetUncompletedTasks()

	th.respondWithJSON(w, http.StatusOK, tasks)
}

func (th *TaskHandler) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	header := mux.Vars(r)["header"]
	task, err := th.uc.GetTask(header)
	if err != nil {
		errResponse := http2.ErrorResponse{
			Message: err.Error(),
			Time:    time.Now(),
		}
		switch {
		case errors.Is(err, domain.ErrTaskNotFound):
			th.respondWithJSON(w, http.StatusOK, errResponse)
		default:
			th.respondWithJSON(w, http.StatusOK, errResponse)
		}
		return
	}

	res := http2.TaskResponse{
		Header:      task.Header,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		CompletedAt: task.CompletedAt,
		CreatedAt:   task.CreatedAt,
	}

	th.respondWithJSON(w, http.StatusOK, res)
}

func (th *TaskHandler) HandleRemoveTask(w http.ResponseWriter, r *http.Request) {
	header := mux.Vars(r)["header"]
	err := th.uc.RemoveTask(header)
	if err != nil {
		errResponse := http2.ErrorResponse{
			Message: err.Error(),
			Time:    time.Now(),
		}

		switch {
		case errors.Is(err, domain.ErrTaskNotFound):
			th.respondWithJSON(w, http.StatusNotFound, errResponse)
		default:
			th.respondWithJSON(w, http.StatusInternalServerError, errResponse)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (th *TaskHandler) respondWithJSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("failed to write http response", err)
		return
	}
}
