package services

import (
	"fmt"
	"go-tasks-api/models"
	"go-tasks-api/repositories"
)

// O service diz, quero algo que seja um TaskRepository
// O service depende de um Repository para funcionar
type TaskService struct {
	repository repositories.TaskRepository
}

// recebe o repository
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

func (s *TaskService) Update(id int, title string, completed bool) (models.Task, bool, error) {

	if title == "" {
		return models.Task{}, false, fmt.Errorf("titulo e obrigatorio")
	}

	task, found := s.repository.GetByID(id)

	if !found {
		return models.Task{}, false, nil
	}

	task.Title = title

	updatedTask, found := s.repository.Update(task)

	return updatedTask, found, nil
}

func (s *TaskService) Patch(
	id int,
	inputTitle *string,
	inputCompleted *bool,
) (models.Task, bool, error) {

	task, found := s.repository.GetByID(id)

	if !found {
		return models.Task{}, false, nil
	}

	if inputTitle != nil {
		if *inputTitle == "" {
			return models.Task{}, false, fmt.Errorf("título é obrigatório")
		}

		task.Title = *inputTitle
	}

	if inputCompleted != nil {
		task.Completed = *inputCompleted
	}

	updatedTask, found := s.repository.Update(task)

	return updatedTask, found, nil
}

func (s *TaskService) Delete(id int) bool {
	return s.repository.Delete(id)
}
