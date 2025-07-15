// Package config implements config structs and
// utility for loading config
package config

import (
	"encoding/json"
	"io"
)

type Config struct {
	API      APIConfig      `json:"api"`
	Database DatabaseConfig `json:"database"`
	Logger   LoggerConfig   `json:"logger"`
}

type APIConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`

	ShutdownIn int `json:"shutdown_in"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoggerConfig struct {
	Environment string `json:"environment"`
}

func LoadConfig(source io.Reader) (*Config, error) {
	var cfg = new(Config)

	decoder := json.NewDecoder(source)
	err := decoder.Decode(cfg)

	return cfg, err
}
