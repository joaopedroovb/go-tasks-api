package models

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var Tasks = []Task{
	{
		ID:	1,
		Title: "Aprender Go",
	},
	{
		ID:    2,
		Title: "Criar API REST",
	},
}