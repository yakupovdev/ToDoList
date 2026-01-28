package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	taskHandler *TaskHandler
}

func NewHTTPServer(taskHandler *TaskHandler) *HTTPServer {
	return &HTTPServer{taskHandler: taskHandler}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
			next.ServeHTTP(w, r)
		})
	})

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
