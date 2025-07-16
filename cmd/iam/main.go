package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"

	"cmd/iam/main.go/internal/config"
)

const _configPath = "CONFIG_PATH"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfgPath := os.Getenv(_configPath)

	if cfgPath == "" {
		return errors.New("not found: CONFIG_PATH")
	}

	if _, err := os.Stat(cfgPath); err != nil {
		return fmt.Errorf("not found %s: %w", cfgPath, err)
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return fmt.Errorf("config file read: %w", err)
	}

	r := bytes.NewReader(data)
	cfg, err := config.LoadConfig(r)
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}
	srv, err := NewServer(cfg)
	if err != nil {
		return fmt.Errorf("server create: %w", err)
	}

	errChan := srv.Run()
	if err = <-errChan; err != nil {
		return fmt.Errorf("server run: %w", err)
	}
	return nil

}
