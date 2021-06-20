package database

import (
	"time"

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

	return &Database{db}
}

func (db *Database) SetConnectionPool(idle, open int, lifetime time.Duration) {
	// set connection pool
	sqlDB, err := db.DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(idle)
	sqlDB.SetMaxOpenConns(open)
	sqlDB.SetConnMaxLifetime(lifetime)
}
