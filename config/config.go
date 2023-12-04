package config

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func LoadAllConfigs(filename string) (configErr error) {

	err := godotenv.Load(filename)
	if err != nil {
		configErr = errors.New("could find the specified .env file")
	}

	loadAppConfig()
	loadDBConfig()

	return configErr
}

func FiberConfig() fiber.Config {

	engine := html.New("./views", ".html")

	return fiber.Config{
		Views: engine,

		ReadTimeout: time.Second * time.Duration(AppConfig.ReadTimeout),
	}
}
