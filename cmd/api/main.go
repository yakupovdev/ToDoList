package main

import (
	"github.com/yakupovdev/ToDoList/internal/delivery/http"
	"github.com/yakupovdev/ToDoList/internal/delivery/http/handler"
	"github.com/yakupovdev/ToDoList/internal/repository"
	"github.com/yakupovdev/ToDoList/internal/usecase"
)

func main() {
	repo := repository.NewTaskRepository()
	uc := usecase.NewTaskUsecase(repo)
	hand := handler.NewTaskHandler(uc)
	server := http.NewHTTPServer(hand)
	err := server.StartServer()
	if err != nil {
		panic(err)
	}
}
