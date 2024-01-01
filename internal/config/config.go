package config

import (
	"log/slog"
	"os"
)

var (
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger() *Logger {
	logger = &Logger{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
	return logger
}
