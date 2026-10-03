package handlers

import (
	"errors"
	"net/http"

	"go-tasks-api/services"
)

func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, services.ErrTaskNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)

	case errors.Is(err, services.ErrInvalidTitle):
		http.Error(w, err.Error(), http.StatusBadRequest)

	default:
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
	}
}
