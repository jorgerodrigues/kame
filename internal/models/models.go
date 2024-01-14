package models

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Model struct {
	User           *UserModel
	URLs           *URLModel
	MonitoringRuns *MonitoringRunModel
}

func New(db *pgxpool.Pool, logger *slog.Logger) *Model {
	return &Model{
		User: &UserModel{
			DB: db,
      logger: logger,
		},
		URLs: &URLModel{
			DB: db,
      logger: logger,
		},
		MonitoringRuns: &MonitoringRunModel{
			DB: db,
      logger: logger,
		},
	}
}
