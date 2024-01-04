package models

import (
	"context"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"-"`
}

type UserModel struct {
	DB *sql.DB
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

	row := m.DB.QueryRowContext(ctx, query, id)
	user := &User{}
	row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	return user
}

func (m *UserModel) FindByEmail(email string) *User {
	query := `SELECT * FROM users WHERE email = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// ok to queryRowContext here since the email has a unique constraint
	row := m.DB.QueryRowContext(ctx, query, email)
	user := &User{}
	row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	return user
}

func (m *UserModel) CreateUser(email, name, plainTextPW string) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	p := &password{
		plaintext: &plainTextPW,
	}
	p.Hash(plainTextPW)

	_, err := m.DB.ExecContext(ctx, query, email, name, p.hash)
	if err != nil {
		// add checks for duplicate email in order to provide better errors
		return err
	}

	return nil
}
