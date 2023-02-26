package cache

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
	modelsapp "todo-api/models/models-app"
	repoapp "todo-api/repository/repo-app"

	"github.com/go-redis/redis"
)

type NewConn struct {
	db *redis.Client
}

func NewRedis(addr string) *NewConn {
	redisAddr := fmt.Sprintf("%s", addr)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	return &NewConn{rdb}
}

func (r *NewConn) GetDataTasks(ctx context.Context, query string) ([]*modelsapp.Task, bool, error) {
	val, err := r.db.Get(query).Result()
	if err == redis.Nil {
		//data db
		page, err := strconv.Atoi(query)
		if err != nil {
			return nil, false, err
		}

		res, err := repoapp.GetTasks(page)
		if err != nil {
			return nil, false, err
		}

		b, err := json.Marshal(res)
		if err != nil {
			return nil, false, err
		}

		err = r.db.Set(query, bytes.NewBuffer(b).Bytes(), time.Hour*3).Err()
		if err != nil {
			return nil, false, err
		}
		// return data not cache
		return res, false, nil
	} else if err != nil {
		return nil, false, errors.New(err.Error() + ", Error in calling redis Client")
	} else {
		data := make([]*modelsapp.Task, 0)
		err := json.Unmarshal(bytes.NewBufferString(val).Bytes(), &data)
		if err != nil {
			return nil, false, err
		}
		return data, true, nil
	}
}

func (r *NewConn) GetDataTask(query string) (*modelsapp.Task, bool, error) {
	val, err := r.db.Get(query).Result()

	if err == redis.Nil {
		id, err := strconv.Atoi(query)
		if err != nil {
			return nil, false, err
		}

		res, err := repoapp.GetTask(uint(id))
		if err != nil {
			return nil, false, err
		}

		b, err := json.Marshal(res)
		if err != nil {
			return nil, false, err
		}

		err = r.db.Set(query, bytes.NewBuffer(b).Bytes(), time.Minute*1).Err()
		if err != nil {
			return nil, false, err
		}

		//return data no cache
		return res, false, nil

	} else if err != nil {
		return nil, false, err
	} else {
		var data *modelsapp.Task
		err := json.Unmarshal(bytes.NewBufferString(val).Bytes(), &data)
		if err != nil {
			return nil, false, err
		}

		return data, true, nil
	}

}
