package models

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jorgerodrigues/upkame/internal/tokens"
)

type Token struct {
	Token     string `json:"token"`
	User      User   `json:"user_id"`
	ExipresAt string `json:"expires_at"`
	CreatedAt string `json:"created_at"`
}

type TokenModel struct {
	DB     *pgxpool.Pool
	logger *slog.Logger
}

func (m *TokenModel) Create(userId string) (string, error) {
	insertQuery := "INSERT INTO tokens (user_id, token, expires_at) VALUES ($1, $2, $3) RETURNING token"
	findUserQuery := "SELECT id, firstname, lastname, email FROM users WHERE id = $1"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRow(ctx, findUserQuery, userId)
	user := &User{}
	err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errors.New("user-not-found")
		} else {
			return "", err
		}
	}

	expiresAt := time.Now().Add(50000 * 24 * time.Hour)
	token, err := tokens.GenerateJWT(userId, []byte("your-secret-key"))
	if err != nil {
		return "", err
	}
	_, err = m.DB.Exec(ctx, insertQuery, userId, token, expiresAt)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (t *TokenModel) Get(token string) string {
	return ""
}

func (t *TokenModel) DeleteToken(tokenId string) error {
	query := "DELETE FROM tokens WHERE id = $1"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := t.DB.Exec(ctx, query, tokenId)
	if err != nil {
		return err
	}

	return nil
}

// validates that the token is in the database and that it is valid with a valid signature
func (t *TokenModel) Validate(token string) (bool, error) {
	query := "SELECT token FROM tokens WHERE token = $1"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := t.DB.QueryRow(ctx, query, token)
	var tokenFromDB string
	err := row.Scan(&tokenFromDB)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, errors.New("token-not-found")
		} else {
			return false, err
		}
	}

	isTheSame := tokenFromDB == token
	if !isTheSame {
		return false, errors.New("token-not-the-same")
	}

	isValid, err := tokens.ValidateJWT(token, []byte("your-secret-key"))
	if err != nil {
		return false, err
	}
	if !isValid {
		return false, errors.New("token-not-valid")
	}

	return true, nil
}
