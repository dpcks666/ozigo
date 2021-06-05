package database

import (
	"ozigo/app/models"
)

func (db *Database) MigrateModels() error {
	err := db.AutoMigrate(
		&models.Role{},
		&models.User{},
	)
	return err
}
