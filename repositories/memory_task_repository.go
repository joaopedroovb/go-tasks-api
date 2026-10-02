package repositories

import "go-tasks-api/models"

type MemoryTaskRepository struct {
	tasks []models.Task
}

func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{
		tasks: []models.Task{
			{
				ID:        1,
				Title:     "Aprender Go",
				Completed: false,
			},
			{
				ID:        2,
				Title:     "Criar API REST",
				Completed: false,
			},
		},
	}
}

func (r *MemoryTaskRepository) GetAll() []models.Task {
	return r.tasks
}

func (r *MemoryTaskRepository) GetByID(id int) (models.Task, bool) {

	for _, task := range r.tasks {

		if task.ID == id {
			return task, true
		}
	}

	return models.Task{}, false
}

func (r *MemoryTaskRepository) Create(task models.Task) models.Task {

	task.ID = len(r.tasks) + 1

	r.tasks = append(r.tasks, task)

	return task
}

func (r *MemoryTaskRepository) Update(task models.Task) (models.Task, bool) {

	for i, currentTask := range r.tasks {

		if currentTask.ID == task.ID {

			r.tasks[i] = task

			return task, true
		}
	}

	return models.Task{}, false
}

func (r *MemoryTaskRepository) Delete(id int) bool {

	var newTasks []models.Task
	encontrou := false

	for _, item := range r.tasks {

		if item.ID == id {
			encontrou = true
			continue
		}

		newTasks = append(newTasks, item)

	}

	if !encontrou {
		return false
	}

	r.tasks = newTasks

	return true
}
