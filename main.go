package main

import (
	"github.com/yakupovdev/ToDoList/internal/handler"
	"github.com/yakupovdev/ToDoList/internal/repository"
	"github.com/yakupovdev/ToDoList/internal/router"
	"github.com/yakupovdev/ToDoList/internal/storage"
	"github.com/yakupovdev/ToDoList/internal/usecase"
)

func main() {
	store := storage.NewStorage()
	repo := repository.NewRepositoryStorage(store)
	uc := usecase.NewTaskUsecase(repo)
	hand := handler.NewTaskHandler(uc)
	server := router.NewHTTPServer(hand)
	err := server.StartServer()
	if err != nil {
		panic(err)
	}

}
