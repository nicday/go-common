package logger

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/nicday/go-common/env"
)

func init() {
	InitLogger()
}

// InitLogger initializes the logger by setting the log level to the env var LOG_LEVEL, or defaulting to `info`.
func InitLogger() {
	// If running in the production environment, output the logs as JSON format for parsing by Logstash.
	if env.IsProd() {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	logrus.SetOutput(os.Stdout)

	level, err := logrus.ParseLevel(env.GetString("LOG_LEVEL", "info"))
	// No need to handle the error here, just don't update the log level
	if err == nil {
		logrus.SetLevel(level)
	}

	logrus.Infof("Log level: %s", logrus.GetLevel().String())
}
