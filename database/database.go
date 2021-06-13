package database

import (
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func New(config gorm.Dialector) *Database {
	// open gorm database
	db, err := gorm.Open(config, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// connect gorm database
	_, err = db.DB()
	if err != nil {
		panic(err)
	}
	return &Database{db}
}
