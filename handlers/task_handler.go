package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go-tasks-api/services"
)

func TasksHandler(
	service *services.TaskService,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		parts := strings.Split(r.URL.Path, "/")

		// /tasks
		if len(parts) == 2 {

			switch r.Method {

			case http.MethodGet:
				getTasks(w, service)

			case http.MethodPost:
				createTask(w, r, service)

			default:
				http.Error(
					w,
					"Método não permitido",
					http.StatusMethodNotAllowed,
				)
			}

			return
		}

		// /tasks/:id
		if len(parts) == 3 {

			id, err := strconv.Atoi(parts[2])

			if err != nil {
				http.Error(
					w,
					"ID inválido",
					http.StatusBadRequest,
				)
				return
			}

			switch r.Method {

			case http.MethodGet:
				getTask(w, service, id)

			case http.MethodPut:
				updateTask(w, r, service, id)
				//w -> como eu respondo ao cliente?
				//r -> o que o cliente me enviou?
				//qual metodo ele enviou
				//r.URL.Path -> qual caminho solicitado ->/tasks/10
				//r.URL.Query() -> quais parms vieram na URL -> /tasks?page=2&limit=10 -> page := r.URL.Query().Get("page")
				//r.Header.Get() -> ler um header da requisicao -> r.Header.Get("Authorization")
				//r.Body -> ler o corpo enviado pelo cliente
				//service -> busca no sistema
				//id -> identifica

			case http.MethodDelete:
				deleteTask(w, service, id)

			default:
				http.Error(
					w,
					"Método não permitido",
					http.StatusMethodNotAllowed,
				)
			}

			return
		}

		http.Error(
			w,
			"Rota não encontrada",
			http.StatusNotFound,
		)
	}
}

func getTasks(
	w http.ResponseWriter,
	service *services.TaskService,
) {

	tasks := service.GetAll()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)
}

func getTask(
	w http.ResponseWriter,
	service *services.TaskService,
	id int,
) {
	task, found := service.GetByID(id)

	if !found {
		http.Error(
			w,
			"Tarefa nao encontrada",
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(task)

}

func createTask(
	w http.ResponseWriter,
	r *http.Request,
	service *services.TaskService,
) {

	var input struct {
		Title string `json:"title"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(
			w, "JSON inválido",
			http.StatusBadRequest,
		)
		return
	}

	task, err := service.Create(input.Title)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(task)
}

func updateTask(
	w http.ResponseWriter,
	r *http.Request,
	service *services.TaskService,
	id int,
) {

	var input struct {
		Title string `json:"title"`
	}

	// Pega o JSON da requisição
	decoder := json.NewDecoder(r.Body)
	// Transforma o JSON em uma struct input
	err := decoder.Decode(&input)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	task, found, err := service.Update(id, input.Title)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	if !found {
		http.Error(
			w,
			"Tarefa não encontrada",
			http.StatusNotFound,
		)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(task)
}

func deleteTask(
	w http.ResponseWriter,
	service *services.TaskService,
	id int,
) {

	found := service.Delete(id)

	if !found {
		http.Error(
			w,
			"Tarefa não encontrada",
			http.StatusNotFound,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
