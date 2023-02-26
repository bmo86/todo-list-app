package database

import (
	modelsapp "todo-api/models/models-app"
)

func (i *instacePostgres) CreateTask(task *modelsapp.Task) (uint, error) {
	err := i.db.Create(&task)
	return task.ID, err.Error
}

func (i *instacePostgres) DeleteTask(id uint) error {
	var task *modelsapp.Task
	if err := i.db.Delete(&task, id).Error; err != nil {
		return err
	}
	return nil
}

func (i *instacePostgres) UpdateTask(id uint, task *modelsapp.Task) (uint, error) {
	data := map[string]interface{}{
		"title":       task.Title,
		"description": task.Description,
		"image":       task.Image,
		"update_at":   task.UpdatedAt,
		"status":      task.Status,
	}

	err := i.db.Table("tasks").Where("idTask = ?", task.ID).UpdateColumns(data)
	if err.Error != nil {
		return 0, err.Error
	}
	return task.ID, nil
}

func (i *instacePostgres) GetTasks(page int) ([]*modelsapp.Task, error) {
	limit := page * 5
	offset := page

	var tasks []*modelsapp.Task
	err := i.db.Offset(offset).Limit(limit).Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (i *instacePostgres) GetTask(id uint) (*modelsapp.Task, error) {
	var task modelsapp.Task
	err := i.db.First(&task, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &task, nil
}
