package repoapp

import modelsapp "todo-api/models/models-app"

type App interface {
	CreateTask(task *modelsapp.Task) (uint, error)
	DeleteTask(id uint) error
	UpdateTask(id uint, task *modelsapp.Task) (uint, error)
	GetTasks() ([]*modelsapp.Task, error)
	GetTask(id uint) (*modelsapp.Task, error)
}

var repo App

func SetRepoApp(r App) {
	repo = r
}

func CreateTask(task *modelsapp.Task) (uint, error) {
	return repo.CreateTask(task)
}

func DeleteTask(id uint) error {
	return repo.DeleteTask(id)
}

func UpdateTask(id uint, task *modelsapp.Task) (uint, error) {
	return repo.UpdateTask(id, task)
}

func GetTasks() ([]*modelsapp.Task, error) {
	return repo.GetTasks()
}

func GetTask(id uint) (*modelsapp.Task, error) {
	return repo.GetTask(id)
}
