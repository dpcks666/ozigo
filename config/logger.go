package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func (config *Config) GetLoggerConfig() zapcore.Core {
	config.SetDefault("LOGGER_TYPE", "console")
	config.SetDefault("LOGGER_FILENAME", "./storage/logs/")

	// Set log level
	var logLevel zapcore.Level
	if config.GetBool("APP_DEBUG") {
		logLevel = zap.DebugLevel
	} else {
		logLevel = zap.InfoLevel
	}

	// Set writer
	var logWriter zapcore.WriteSyncer
	if config.GetString("LOGGER_TYPE") == "file" {
		logWriter = zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.GetString("LOGGER_FILENAME"),
			MaxSize:    config.GetInt("LOGGER_MAX_SIZE"), // megabytes
			MaxAge:     config.GetInt("LOGGER_MAX_AGE"),  // days
			MaxBackups: config.GetInt("LOGGER_MAX_BACKUPS"),
			LocalTime:  config.GetBool("LOGGER_LOCAL_TIME"),
		})
	} else {
		w, _, err := zap.Open("stderr")
		if err != nil {
			panic(err)
		}
		logWriter = w
	}

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		logWriter,
		zap.NewAtomicLevelAt(logLevel),
	)
}
