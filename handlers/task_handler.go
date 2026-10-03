package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go-tasks-api/services"
)

// O cliente ira mandar so o title no body
// DTO -> o que a API recebe
type CreateTaskInput struct {
	Title string `json:"title"`
}

// Sem o ponteiro, se usuario tentasse atualizar sem o Completed
// o go iria tratar ele como false
type UpdateTaskInput struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// agora com ponteiro ele so altera o que chegar com valor
// se tiver sem mantem o valor
type PatchTaskInput struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

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

			case http.MethodPatch:
				patchTask(w, r, service, id)

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
	task, err := service.GetByID(id)

	if err != nil {
		writeError(w, err)
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

	var input CreateTaskInput

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
		writeError(w, err)
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

	var input UpdateTaskInput

	// Pega o JSON da requisição
	decoder := json.NewDecoder(r.Body)
	// Transforma o JSON em uma struct input
	err := decoder.Decode(&input)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	task, err := service.Update(
		id,
		input.Title,
		input.Completed,
	)

	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(task)
}

func patchTask(
	w http.ResponseWriter,
	r *http.Request,
	service *services.TaskService,
	id int,
) {
	var input PatchTaskInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	task, err := service.Patch(
		id,
		input.Title,
		input.Completed,
	)

	if err != nil {
		writeError(w, err)
		return
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

	err := service.Delete(id)

	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
