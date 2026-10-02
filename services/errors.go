package services

import "errors"

var (
	ErrTaskNotFound = errors.New("tarefa nao encontrada")
	ErrInvalidTitle = errors.New("titulo e obrigatorio")
)
