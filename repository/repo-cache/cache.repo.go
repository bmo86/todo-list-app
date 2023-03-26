package repocache

import (
	"context"
	modelsapp "todo-api/models/models-app"
	modelsusr "todo-api/models/models-usr"
)

type Cache interface {
	GetDataTasks(ctx context.Context, query string) ([]*modelsapp.Task, bool, error)
	GetDataTask(query string) (*modelsapp.Task, bool, error)
	GetUser_ID(ctx context.Context, query string) (*modelsusr.GetUsr_ID, bool, error)
	GetUser_Email(query string) (*modelsusr.GetUsr_Email, bool, error)
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

func GetUser_ID(ctx context.Context, query string) (*modelsusr.GetUsr_ID, bool, error) {
	return cache.GetUser_ID(ctx, query)
}

func GetUser_Email(query string) (*modelsusr.GetUsr_Email, bool, error) {
	return cache.GetUser_Email(query)
}
