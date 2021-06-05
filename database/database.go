package database

import (
	"log"

	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func New(config gorm.Dialector) *Database {
	db, err := gorm.Open(config, &gorm.Config{})
	if err != nil {
		log.Println("failed to connect to database:", err.Error())
	}
	return &Database{db}
}
