package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yakupovdev/ToDoList/internal/delivery/http/handler"
	"github.com/yakupovdev/ToDoList/internal/delivery/http/middleware"
)

type HTTPServer struct {
	taskHandler *handler.TaskHandler
}

func NewHTTPServer(taskHandler *handler.TaskHandler) *HTTPServer {
	return &HTTPServer{taskHandler: taskHandler}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Use(middleware.LoggerMiddleware)

	router.Path("/tasks").Methods("POST").HandlerFunc(s.taskHandler.HandleAddTask)
	router.Path("/tasks").Queries("completed", "false").Methods("GET").HandlerFunc(s.taskHandler.HandleGetUncompletedTasks)
	router.Path("/tasks/{header}").Methods("PATCH").HandlerFunc(s.taskHandler.HandleChangeCompleteStatusTask)
	router.Path("/tasks/{header}").Methods("DELETE").HandlerFunc(s.taskHandler.HandleRemoveTask)
	router.Path("/tasks/{header}").Methods("GET").HandlerFunc(s.taskHandler.HandleGetTask)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.taskHandler.HandleGetAllTasks)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return err
	}
	return nil
}
