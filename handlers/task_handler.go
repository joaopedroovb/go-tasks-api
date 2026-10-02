package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go-tasks-api/models"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")

	// /tasks
	if len(parts) == 2 {

		switch r.Method {
		case http.MethodGet:
			getTasks(w, r)

		case http.MethodPost:
			createTask(w, r)
		
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}

		return
	}

	// /tasks/:id
	if len(parts) == 3 {

		id, err := strconv.Atoi(parts[2])

		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			getTask(w, r, id)

		case http.MethodPut:
			updateTask(w, r, id)

		case http.MethodDelete:
			deleteTask(w, r, id)

		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}

		return
	}

	http.Error(w, "Rota não encontrada", http.StatusNotFound)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.Tasks)
}

func getTask(w http.ResponseWriter, r *http.Request, id int) {
	for _, task := range models.Tasks {
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
	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		http.Error(w, "Título é obrigatório", http.StatusBadRequest)
		return
	}

	task.ID = len(models.Tasks) + 1
	models.Tasks = append(models.Tasks, task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request, id int) {
	var updatedTask models.Task

	// Pega o JSON da requisição 
	decoder := json.NewDecoder(r.Body)
	// Transforma o JSON em uma struct Task
	err := decoder.Decode(&updatedTask)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if updatedTask.Title == ""{
		http.Error(w, "Título é obrigatório", http.StatusBadRequest)
		return
	}

	for i, task := range models.Tasks {
		if task.ID == id {
			models.Tasks[i].Title = updatedTask.Title

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(models.Tasks[i])
			return
		}
	}

	http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
}

func deleteTask(w http.ResponseWriter, r *http.Request, id int) {

	var newTasks []models.Task
	encontrou := false

	for _, task := range models.Tasks {

		if task.ID == id {
			encontrou = true
			continue
		}

		newTasks = append(newTasks, task)
	}

	if !encontrou {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	models.Tasks = newTasks

	w.WriteHeader(http.StatusNoContent)
}
