package config

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	jaeger "github.com/uber/jaeger-client-go/config"
)

func (config *Config) GetTracerConfig() *jaeger.Configuration {
	config.SetDefault("JAEGER_SERVICE_NAME", "fiber/"+fiber.Version)
	config.SetDefault("JAEGER_SAMPLER_TYPE", "const")
	config.SetDefault("JAEGER_SAMPLER_PARAM", 1)
	config.SetDefault("JAEGER_REPORTER_LOG_SPANS", true)
	config.SetDefault("JAEGER_REPORTER_FLUSH_INTERVAL", "1s")

	defcfg := jaeger.Configuration{
		ServiceName: config.GetString("JAEGER_SERVICE_NAME"),
		Sampler: &jaeger.SamplerConfig{
			Type:  config.GetString("JAEGER_SAMPLER_TYPE"),
			Param: config.GetFloat64("JAEGER_SAMPLER_PARAM"),
		},
		Reporter: &jaeger.ReporterConfig{
			LogSpans:            config.GetBool("JAEGER_REPORTER_LOG_SPANS"),
			BufferFlushInterval: config.GetDuration("JAEGER_REPORTER_FLUSH_INTERVAL"),
		},
	}
	cfg, err := defcfg.FromEnv()
	if err != nil {
		fmt.Println("Could not parse Jaeger env vars: ", err.Error())
	}

	return cfg
}
