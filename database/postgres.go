package database

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type instacePostgres struct {
	db *gorm.DB
}

func NewConnectionDB(url string) (*instacePostgres, error) {
	db, err := gorm.Open(postgres.Open(url))
	if err != nil {
		return nil, err
	}

	return &instacePostgres{db}, err
}
