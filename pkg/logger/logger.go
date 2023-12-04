package logger

import (
	"os"

	"github.com/rakesh-gupta29/sqlite-golang/config"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

var logger = &Logger{}

// Setup Logger settings
func SetUpLogger() {
	logger = &Logger{logrus.New()}
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetOutput(os.Stdout)

	if config.AppConfig.DebugMode {
		logger.SetLevel(logrus.DebugLevel)
	}
}

func GetLogger() *Logger {
	return logger
}
