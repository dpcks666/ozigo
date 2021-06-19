package database

import (
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func New(dialector gorm.Dialector) *Database {
	// open gorm database
	db, err := gorm.Open(dialector, &gorm.Config{})
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
