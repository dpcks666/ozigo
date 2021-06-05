package config

import (
	"github.com/gofiber/fiber/v2"
)

func (config *Config) GetFiberConfig() fiber.Config {
	return fiber.Config{
		Prefork:               config.GetBool("FIBER_PREFORK"),
		ServerHeader:          config.GetString("APP_NAME"),
		UnescapePath:          config.GetBool("FIBER_UNESCAPEPATH"),
		BodyLimit:             config.GetInt("FIBER_BODYLIMIT"),
		Concurrency:           config.GetInt("FIBER_CONCURRENCY"),
		Views:                 config.getFiberViewsEngine(),
		ProxyHeader:           config.GetString("FIBER_PROXYHEADER"),
		DisableStartupMessage: !config.GetBool("APP_DEBUG"),
	}
}
