package models

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Model struct {
	User                 *UserModel
	URLs                 *URLModel
	MonitoringRuns       *MonitoringRunModel
	Tokens               *TokenModel
	Organizations        *OrganizationModel
	UsersOnOrganizations *UsersOnOrganizationModel
}

func New(db *pgxpool.Pool, logger *slog.Logger) *Model {
	return &Model{
		User: &UserModel{
			DB:     db,
			logger: logger,
		},
		URLs: &URLModel{
			DB:     db,
			logger: logger,
		},
		MonitoringRuns: &MonitoringRunModel{
			DB:     db,
			logger: logger,
		},
		Tokens: &TokenModel{
			DB:     db,
			logger: logger,
		},
		Organizations: &OrganizationModel{
			DB:     db,
			logger: logger,
		},
		UsersOnOrganizations: &UsersOnOrganizationModel{
			DB:     db,
			logger: logger,
		},
	}
}
