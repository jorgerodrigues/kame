package models

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersOnOrganizations struct {
	ID       string    `json:"id"`
	UserID   string    `json:"user_id"`
	OrgId    string    `json:"organization_id"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
}

type UsersOnOrganizationModel struct {
	DB     *pgxpool.Pool
	logger *slog.Logger
}

func (m *UsersOnOrganizationModel) Create(userID, organization, role string) (string, error) {
	insertUserOnOrganization := `INSERT INTO users_on_organizations (user_id, organization_id, role) VALUES ($1, $2, $3) RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRow(ctx, insertUserOnOrganization, userID, organization, role)

	uoo := &UsersOnOrganizations{}
	err := row.Scan(&uoo.ID)
	if err != nil {
		m.logger.Error("Error scanning row", err)
		return "", err
	}

	return uoo.ID, nil
}

func (m *UsersOnOrganizationModel) FindByOrgRoleAndUser(userID, orgId, role string) (*UsersOnOrganizations, error) {
	findExistingUserOnOrganization := `SELECT id, user_id, organization_id, role, created_at, updated_at FROM users_on_organizations WHERE user_id = $1 AND organization_id = $2 AND role = $3`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRow(ctx, findExistingUserOnOrganization, userID, orgId, role)
	userOnOrg := &UsersOnOrganizations{}

	err := row.Scan(&userOnOrg.ID, &userOnOrg.UserID, &userOnOrg.OrgId, &userOnOrg.CreateAt, &userOnOrg.UpdateAt)
	if err != nil {
		m.logger.Error("Error scanning row", err)
		return nil, err
	}
	return userOnOrg, nil
}
