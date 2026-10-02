package services

import (
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
		return models.Task{}, ErrInvalidTitle
	}

	task := models.Task{
		Title:     title,
		Completed: false,
	}

	return s.repository.Create(task), nil
}

func (s *TaskService) Update(
	id int,
	title string,
	completed bool,
) (models.Task, error) {

	if title == "" {
		return models.Task{}, ErrInvalidTitle
	}

	task, found := s.repository.GetByID(id)

	if !found {
		return models.Task{}, ErrTaskNotFound
	}

	task.Title = title
	task.Completed = completed

	return s.repository.Update(task), nil
}

func (s *TaskService) Patch(
	id int,
	inputTitle *string,
	inputCompleted *bool,
) (models.Task, error) {

	task, found := s.repository.GetByID(id)

	if !found {
		return models.Task{}, ErrTaskNotFound
	}

	if inputTitle != nil {
		if *inputTitle == "" {
			return models.Task{}, ErrInvalidTitle
		}

		task.Title = *inputTitle
	}

	if inputCompleted != nil {
		task.Completed = *inputCompleted
	}

	return s.repository.Update(task), nil
}

func (s *TaskService) Delete(id int) bool {
	return s.repository.Delete(id)
}
