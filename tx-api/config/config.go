package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database Database `yaml:"database"`
}
type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Schema   string `yaml:"schema"`
}

var conf *Config

func Load() error {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local" // default
	}

	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return fmt.Errorf("failed to read config.yaml: %w", err)
	}

	all := make(map[string]Config)
	if err := yaml.Unmarshal(data, &all); err != nil {
		return fmt.Errorf("failed to unmarshal config.yaml: %w", err)
	}

	cfg, ok := all[env]
	if !ok {
		return fmt.Errorf("environment %q not found in config.yaml", env)
	}

	conf = &cfg
	return nil
}

func Get() (*Config, error) {
	if conf == nil {
		if err := Load(); err != nil {
			return nil, fmt.Errorf("config load failed: %v", err)
		}
	}
	return conf, nil
}
