package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/dynamodb"
	"github.com/gofiber/storage/memcache"
	"github.com/gofiber/storage/memory"
	"github.com/gofiber/storage/mongodb"
	"github.com/gofiber/storage/mysql"
	"github.com/gofiber/storage/postgres"
	"github.com/gofiber/storage/redis"
	"github.com/gofiber/storage/sqlite3"
)

func (config *Config) GetSessionConfig() session.Config {
	var storage fiber.Storage
	switch config.GetString("SESSION_PROVIDER") {
	case "dynamodb":
		storage = dynamodb.New(dynamodb.Config{
			Table: config.GetString("SESSION_TABLENAME"),
		})
	case "memcache":
		storage = memcache.New(memcache.Config{
			Servers: config.GetString("SESSION_HOST") + ":" + config.GetString("SESSION_PORT"),
		})
	case "mongodb":
		storage = mongodb.New(mongodb.Config{
			Host:       config.GetString("SESSION_HOST"),
			Port:       config.GetInt("SESSION_PORT"),
			Database:   config.GetString("SESSION_DATABASE"),
			Collection: config.GetString("SESSION_TABLENAME"),
		})
	case "mysql":
		storage = mysql.New(mysql.Config{
			Host:     config.GetString("SESSION_HOST"),
			Port:     config.GetInt("SESSION_PORT"),
			Username: config.GetString("SESSION_USERNAME"),
			Password: config.GetString("SESSION_PASSWORD"),
			Database: config.GetString("SESSION_DATABASE"),
			Table:    config.GetString("SESSION_TABLENAME"),
		})
	case "postgres":
		storage = postgres.New(postgres.Config{
			Host:     config.GetString("SESSION_HOST"),
			Port:     config.GetInt("SESSION_PORT"),
			Database: config.GetString("SESSION_DATABASE"),
			Table:    config.GetString("SESSION_TABLENAME"),
		})
	case "redis":
		storage = redis.New(redis.Config{
			Host:     config.GetString("SESSION_HOST"),
			Port:     config.GetInt("SESSION_PORT"),
			Username: config.GetString("SESSION_USERNAME"),
			Password: config.GetString("SESSION_PASSWORD"),
			Database: config.GetInt("SESSION_DATABASE"),
		})
	case "sqlite3":
		storage = sqlite3.New(sqlite3.Config{
			Database: config.GetString("SESSION_DATABASE"),
			Table:    config.GetString("SESSION_TABLENAME"),
		})
	default:
		storage = memory.New()
	}
	return session.Config{
		Expiration:     config.GetDuration("SESSION_EXPIRATION"),
		Storage:        storage,
		CookieHTTPOnly: true,
	}
}
