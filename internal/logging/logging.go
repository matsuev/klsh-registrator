package logging

import (
	"go.uber.org/zap"
)

// Config ...
type Config struct {
	IsDebug bool `yaml:"is_debug" env-default:"false"`
}

// DefaultConfig function
func DefaultConfig() *Config {
	return &Config{
		IsDebug: false,
	}
}

// Logger ...
type Logger struct {
	*zap.SugaredLogger
}

// New function
func New(cfg *Config) (*Logger, error) {
	var logger *zap.Logger
	var err error

	if cfg == nil {
		cfg = DefaultConfig()
	}

	if cfg.IsDebug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	return &Logger{
		logger.Sugar(),
	}, nil
}
