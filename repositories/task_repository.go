package repositories

import "go-tasks-api/models"

type TaskRepository interface {
	GetAll() []models.Task
	GetByID(id int) (models.Task, bool)
	Create(task models.Task) models.Task
	Update(task models.Task) (models.Task, bool)
	Delete(id int) bool
}
