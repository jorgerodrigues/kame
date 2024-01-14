package models

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string         `json:"id"`
	Firstname sql.NullString `json:"firstname"`
	Lastname  sql.NullString `json:"lastname"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"-"`
}

type UserModel struct {
	DB     *pgxpool.Pool
	logger *slog.Logger
}

type password struct {
	plaintext *string
	hash      []byte
}

func (p *password) Hash(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.plaintext = &plaintextPassword
	p.hash = hash

	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *UserModel) FindById(id string) *User {
	query := `SELECT * FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRow(ctx, query, id)
	user := &User{}
	row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.CreatedAt)

	return user
}

func (m *UserModel) FindByEmail(email string) (*User, error) {
	query := `SELECT id, firstname, lastname, email, password_hash, created_at FROM users WHERE email = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// ok to queryRowContext here since the email has a unique constraint
	var user User
	err := m.DB.QueryRow(ctx, query, email).Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UserModel) CreateUser(email, firstname, lastname, plainTextPW string) error {
	query := `INSERT INTO users (firstname, lastname, email, password_hash) VALUES ($1, $2, $3 , $4)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	p := &password{
		plaintext: &plainTextPW,
	}
	p.Hash(plainTextPW)

	_, err := m.DB.Exec(ctx, query, firstname, lastname, email, p.hash)
	if err != nil {
		// add checks for duplicate email in order to provide better errors
		return err
	}

	return nil
}
