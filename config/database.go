package config

import (
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func (config *Config) GetDatabaseDialector() (dialector gorm.Dialector) {
	switch config.GetString("DB_DRIVER") {
	case "mysql":
		dsn := config.GetString("DB_USERNAME") + ":" + config.GetString("DB_PASSWORD") + "@tcp(" + config.GetString("DB_HOST") + ":" + strconv.Itoa(config.GetInt("DB_PORT")) + ")/" + config.GetString("DB_DATABASE") + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=UTC"
		dialector = mysql.Open(dsn)
	case "postgresql", "postgres":
		dsn := "user=" + config.GetString("DB_USERNAME") + " password=" + config.GetString("DB_PASSWORD") + " dbname=" + config.GetString("DB_DATABASE") + " host=" + config.GetString("DB_HOST") + " port=" + strconv.Itoa(config.GetInt("DB_PORT")) + " TimeZone=UTC"
		dialector = postgres.Open(dsn)
	case "sqlserver", "mssql":
		dsn := "sqlserver://" + config.GetString("DB_USERNAME") + ":" + config.GetString("DB_PASSWORD") + "@" + config.GetString("DB_HOST") + ":" + strconv.Itoa(config.GetInt("DB_PORT")) + "?database=" + config.GetString("DB_DATABASE")
		dialector = sqlserver.Open(dsn)
	}
	return
}
