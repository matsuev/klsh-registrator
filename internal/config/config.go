package config

import (
	"flag"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/matsuev/klsh-registrator/internal/logging"
	"github.com/matsuev/klsh-registrator/internal/router"
	"github.com/matsuev/klsh-registrator/internal/server"
)

const (
	DEFAULT_CONFIG_PATH = "./configs/config.yml"
)

// Config ...
type Config struct {
	Shutdown time.Duration   `yaml:"shutdown" env-default:"10s"`
	Logger   *logging.Config `yaml:"logger"`
	Server   *server.Config  `yaml:"server"`
	Router   *router.Config  `yaml:"router"`
}

// DefaultConfig function
func DefaultConfig() *Config {
	return &Config{
		Shutdown: 10 * time.Second,
		Logger:   logging.DefaultConfig(),
		Server:   server.DefaultConfig(),
		Router:   router.DefaultConfig(),
	}
}

// New function
func New() (*Config, error) {
	cfgPath := flag.String("c", "", "Path to <config.yml> file. Default <./configs/config.yml>")
	flag.Parse()

	if *cfgPath == "" {
		*cfgPath = DEFAULT_CONFIG_PATH
	}

	cfg := DefaultConfig()

	if err := cleanenv.ReadConfig(*cfgPath, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
