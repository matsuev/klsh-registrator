package server

import (
	"fmt"
	"net/http"
	"time"
)

const (
	DEFAULT_SERVER_ADDR      = "127.0.0.1"
	DEFAULT_SERVER_PORT      = "8080"
	DEFAULT_READ_TIMEOUT     = 5 * time.Second
	DEFAULT_WRITE_TIMEOUT    = 5 * time.Second
	DEFAULT_MAX_HEADER_BYTES = 1 << 20 // 1048576
)

// Config ...
type Config struct {
	Addr           string        `yaml:"addr" env-default:"127.0.0.1"`
	Port           string        `yaml:"port" env-default:"8080"`
	ReadTimeout    time.Duration `yaml:"read_timeout" env-default:"5s"`
	WriteTimeout   time.Duration `yaml:"write_timeout" env-default:"5s"`
	MaxHeaderBytes int           `yaml:"max_bytes" env-default:"1048576"`
}

// DefaultConfig function
func DefaultConfig() *Config {
	return &Config{
		Addr:           DEFAULT_SERVER_ADDR,
		Port:           DEFAULT_SERVER_PORT,
		ReadTimeout:    DEFAULT_READ_TIMEOUT,
		WriteTimeout:   DEFAULT_WRITE_TIMEOUT,
		MaxHeaderBytes: DEFAULT_MAX_HEADER_BYTES,
	}
}

// New function
func New(cfg *Config, handler http.Handler) *http.Server {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	return &http.Server{
		Addr:           fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port),
		Handler:        handler,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
	}
}
