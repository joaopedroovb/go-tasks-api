package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	http.HandleFunc("/tasks", tasksHandler)

	fmt.Println("Servidor rodando na porta 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		getTasks(w, r)

	case http.MethodPost:
		createTask(w, r)

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []Task{
		{
			ID:    1,
			Title: "Aprender Go",
		},
		{
			ID:    2,
			Title: "Criar API REST",
		},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	task.ID = 3

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(task)
}
