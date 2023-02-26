package modelscache

import modelsapp "todo-api/models/models-app"

type CacheResponse struct {
	Cache bool              `json:"cache"`
	Data  []*modelsapp.Task `json:"data"`
}

type CacheResponseOnlyOne struct {
	Cache bool            `json:"cache"`
	Data  *modelsapp.Task `json:"data"`
}
