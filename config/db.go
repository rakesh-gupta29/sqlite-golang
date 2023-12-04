package config

import (
	"errors"
	"os"
)

type DatabaseConfig struct {
	URL string
}

var DBConfig DatabaseConfig

func loadDBConfig() error {
	url := os.Getenv("MONGO_URI")
	if url == "" {
		return errors.New("invalid database URL")
	}
	DBConfig.URL = url
	return nil
}
