package modelsapp

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	User_id     uint      `json:"usr_id"`
	Check       bool      `json:"check"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       []byte    `json:"image"`
	DateFinish  time.Time `json:"date_finish"`
	Status      bool      `json:"status"`
}

type Request_Task struct {
	User_id     uint      `json:"usr_id"`
	Check       bool      `json:"check"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	DateFinish  time.Time `json:"date_finish"`
	Status      bool      `json:"status"`
}
