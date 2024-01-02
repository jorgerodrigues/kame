package config

import (
	"log/slog"
	"net/http"
)

type Logger struct {
  logger *slog.Logger
}

func (l *Logger) LogError(r* http.Request, err error) {
	l.logger.Error(err.Error())
}

func (l *Logger) LogWarning(r* http.Request, warn string) {
	l.logger.Warn(warn)
}
func (l *Logger) LogInfo(r* http.Request, info string) {
  l.logger.Info(info)
}


