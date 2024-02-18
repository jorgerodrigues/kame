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

func (m *OrganizationModel) Create(name string) (*Organization, error) {
	insertOrganization := `INSERT INTO organizations (name) VALUES ($1) RETURNING id, name, created_at`
	org := &Organization{}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRow(ctx, insertOrganization, name)
	err := row.Scan(&org.ID, &org.Name, &org.CreateAt)

	if err != nil {
		m.logger.Error("Error inserting organization", err)
		return nil, err
	}

	return org, nil
}

func (m *OrganizationModel) FindByName(name string) (*Organization, error) {
	checkOrgExists := `SELECT id, name, created_at FROM organizations WHERE name = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRow(ctx, checkOrgExists, name)
	org := &Organization{}

	err := row.Scan(&org.ID, &org.Name, &org.CreateAt)

	if err != nil && err.Error() == "no rows in result set" {
		return nil, nil
	}

	if err != nil && err.Error() != "no rows in result set" {
		m.logger.Error("Error scanning row", err)
		return nil, err
	}

	return org, nil

}

func (m *OrganizationModel) FindById(id string) (*Organization, error) {
	checkOrgExists := `SELECT id, name, created_at FROM organizations WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRow(ctx, checkOrgExists, id)
	org := &Organization{}

	err := row.Scan(&org.ID, &org.Name, &org.CreateAt)

	if err != nil && err.Error() == "no rows in result set" {
		return nil, nil
	}

	if err != nil && err.Error() != "no rows in result set" {
		m.logger.Error("Error scanning row", err)
		return nil, err
	}

	return org, nil
}

func (m *OrganizationModel) AddUserToOrganization(userID, orgID, role string) error {
	insertUserOnOrganization := `INSERT INTO users_on_organizations (user_id, organization_id, role) VALUES ($1, $2, $3)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.Exec(ctx, insertUserOnOrganization, userID, orgID, role)
	if err != nil {
		m.logger.Error("Error inserting user on organization", err)
		return err
	}

	return nil
}
