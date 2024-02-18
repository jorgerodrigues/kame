package models

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Organization struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"created_at"`
}

type OrganizationModel struct {
	DB     *pgxpool.Pool
	logger *slog.Logger
}

func (m *OrganizationModel) Create(name string) error {
	insertOrganization := `INSERT INTO organizations (name) VALUES ($1) RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.Exec(ctx, insertOrganization, name)
	if err != nil {
		m.logger.Error("Error inserting organization", err)
		return err
	}

	return nil
}

func (m *OrganizationModel) FindByName(name string) (*Organization, error) {
	checkOrgExists := `SELECT id FROM organizations WHERE name = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRow(ctx, checkOrgExists, name)
	org := &Organization{}

	err := row.Scan(&org.ID)

	if err != nil && err.Error() == "no rows in result set" {
		return nil, nil
	}

	if err != nil && err.Error() != "no rows in result set" {
		m.logger.Error("Error scanning row", err)
		return nil, err
	}

	return org, nil

}
