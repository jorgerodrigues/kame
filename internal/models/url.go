package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type URL struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	OwnerId     string `json:"ownerId"`
	CreatedById string `json:"createdById"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type URLModel struct {
	DB *pgxpool.Pool
}

func (m *URLModel) Create(url string, name string, ownerId, createdById string) error {
	query := `INSERT INTO urls (url, name, owner_id, created_by_id) VALUES ($1, $2, $3, $4) RETURNING id`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.Exec(ctx, query, url, name, ownerId, createdById)

	if err != nil {
		return err
	}
	return nil
}

func (m *URLModel) GetById(id string) (*URL, error) {
	query := `SELECT id, url, name, owner_id, created_by_id FROM urls WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRow(ctx, query, id)

	var url URL
	err := row.Scan(
		&url.ID,
		&url.URL,
		&url.Name,
		&url.OwnerId,
		&url.CreatedById,
	)

	if err != nil {
		return nil, err
	}

	return &url, nil
}
