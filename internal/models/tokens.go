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
func (m *TokenModel) createToken(userId string) (string, error) {
  insertQuery := "INSERT INTO tokens (user_id, token, expires_at) VALUES ($1, $2, $3) RETURNING token"
  findUserQuery := "SELECT id, firstname, lastname, email FROM users WHERE id = $1"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRow(ctx, findUserQuery, userId)
	user := &User{}
  err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email )
  if err != nil {
    if errors.Is(err, pgx.ErrNoRows) {
      return "", errors.New("user-not-found")
    } else {
      return "", err
    }
  }

  expiresAt := time.Now().Add(50000 * 24 * time.Hour)
  token, err := tokens.GenerateJWT(userId)
  if err != nil {
    return "", err
  }
  _, err = m.DB.Exec(ctx, insertQuery, userId, token, expiresAt)
  if err != nil {
    return "", err
  }
	return token, nil
}

func (t *TokenModel) getToken(token string) string {
	return ""
}

func (t *TokenModel) deleteToken(token string) string {
	return ""
}

func (t *TokenModel) validateToken() (bool, error) {
  return true, nil
}
