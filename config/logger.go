package config

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (config *Config) GetLoggerConfig() zap.Config {
	// Set print log level
	var logLevel zapcore.Level
	if config.GetBool("APP_DEBUG") {
		logLevel = zap.DebugLevel
	} else {
		logLevel = zap.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    "function",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return zap.Config{
		Level:            zap.NewAtomicLevelAt(logLevel),
		Development:      config.GetBool("APP_DEBUG"),
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func (config *Config) GetAccessLoggerConfig(skipper func(c *fiber.Ctx) bool) logger.Config {
	return logger.Config{
		Next: skipper,
		Format: `{"time":"${time}","id":"${id}","remote_ip":"${ip}",` +
			`"method":"${method}","status":${status},"host":"${host}","url":"${url}","userAgent":"${ua}",` +
			`"error":"${error}","latency":${latency},"latencyParse"` + "\n",
		TimeFormat:   "2006-01-02T15:04:05.000Z0700",
		TimeInterval: 200 * time.Millisecond,
		Output:       os.Stdout,
	}
}
