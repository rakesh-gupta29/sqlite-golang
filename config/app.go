package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Host           string
	Port           string
	DebugMode      bool
	ReadTimeout    time.Duration
	AllowedDomains []string
}

var AppConfig Config

func loadAppConfig() {
	AppConfig.Host = os.Getenv("HOST")
	if AppConfig.Host == "" {
		AppConfig.Host = "0.0.0.0"
	}

	AppConfig.Port = os.Getenv("PORT")
	if AppConfig.Port == "" {
		AppConfig.Port = "3000"
	}
	AppConfig.AllowedDomains = strings.Split(os.Getenv("ALLOWED_DOMAINS"), ",")
	if AppConfig.Port == "" {
		AppConfig.Port = "3000"
	}
	debugMode, err := strconv.ParseBool(os.Getenv("APP_DEBUG_MODE"))
	if err != nil {
		debugMode = false
	}
	AppConfig.DebugMode = debugMode

	timeout, err := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	if err != nil {
		timeout = 30
	}
	AppConfig.ReadTimeout = time.Second * time.Duration(timeout)
}
