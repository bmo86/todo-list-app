package repocache

import (
	"context"
	modelsapp "todo-api/models/models-app"
)

type Cache interface {
	GetDataTasks(ctx context.Context, query string) ([]*modelsapp.Task, bool, error)
	GetDataTask(query string) (*modelsapp.Task, bool, error)
}

var cache Cache

func SetRepo(c Cache) {
	cache = c
}

func GetDataTasks(ctx context.Context, query string) ([]*modelsapp.Task, bool, error) {
	return cache.GetDataTasks(ctx, query)
}

func GetDataTask(query string) (*modelsapp.Task, bool, error) {
	return cache.GetDataTask(query)
}
