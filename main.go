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
	http.HandleFunc("/tasks", getTasks)

	fmt.Println("Servidor rodando na porta 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
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
