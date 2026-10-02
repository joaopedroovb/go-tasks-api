package services

import (
	"fmt"
	"go-tasks-api/models"
	"go-tasks-api/repositories"
)

type TaskService struct {
	repository repositories.TaskRepository
}

func NewTaskService(repository repositories.TaskRepository) *TaskService {
	return &TaskService{
		repository: repository,
	}
}

func (s *TaskService) GetAll() []models.Task {
	return s.repository.GetAll()
}

func (s *TaskService) GetByID(id int) (models.Task, bool) {
	return s.repository.GetByID(id)
}

func (s *TaskService) Create(title string) (models.Task, error) {

	if title == "" {
		return models.Task{}, fmt.Errorf("titulo e obrigatorio")
	}

	task := models.Task{
		Title:     title,
		Completed: false,
	}

	return s.repository.Create(task), nil
}

func (s *TaskService) Update(id int, title string) (models.Task, bool, error) {

	if title == "" {
		return models.Task{}, false, fmt.Errorf("titulo e obrigatorio")
	}

	task := models.Task{
		ID:        id,
		Title:     title,
		Completed: false,
	}

	updatedTask, found := s.repository.Update(task)

	return updatedTask, found, nil
}

func (s *TaskService) Delete(id int) bool {
	return s.repository.Delete(id)
}
