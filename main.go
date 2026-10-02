package main

import (
	"fmt"
	"net/http"

	"go-tasks-api/handlers"
	"go-tasks-api/repositories"
	"go-tasks-api/services"
)

func main() {

	repository := repositories.NewMemoryTaskRepository()

	service := services.NewTaskService(repository)

	http.HandleFunc(
		"/tasks",
		handlers.TasksHandler(service),
	)

	http.HandleFunc(
		"/tasks/",
		handlers.TasksHandler(service),
	)

	fmt.Println("Servidor rodando na porta 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}
