package main

import (
	"fmt"
	"net/http"

	"go-tasks-api/handlers"
	"go-tasks-api/repositories"
	"go-tasks-api/services"
)

func main() {

	// 1. Cria a implementação concreta do repository
	repository := repositories.NewMemoryTaskRepository()

	// 2. Injeta o repository no service
	service := services.NewTaskService(repository)

	// 3. Injeta o service no handler
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
