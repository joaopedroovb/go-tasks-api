package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"strconv"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var tasks = []Task{
	{
		ID:	1,
		Title: "Aprender Go",
	},
	{
		ID:    2,
		Title: "Criar API REST",
	},
}

func main() {
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", tasksHandler)

	fmt.Println("Servidor rodando na porta 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) == 3 && r.Method == http.MethodGet {
		id, err := strconv.Atoi(parts[2])

		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		getTask(w, r, id)
		return
	}

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
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request, id int) {
	for _, task := range tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
	
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	task.ID = len(tasks) + 1
	tasks = append(tasks, task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(task)
}
