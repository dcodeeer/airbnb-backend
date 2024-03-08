package users

import (
	"api/internal/core"
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type IMailService interface {
	Send(ctx context.Context, to, template string, content map[string]string) error
}

type users struct {
	db          *sqlx.DB
	mailService IMailService
}

func New(db *sqlx.DB) *users {
	return &users{db: db}
}

func (r *users) Create(ctx context.Context, user *core.User) (int, error) {
	var id int
	sql := "INSERT INTO users.users (email, password) VALUES ($1, $2) RETURNING id;"
	err := r.db.QueryRowx(sql, user.Email, user.Password).Scan(&id)
	return id, err
}

func (r *users) GetOneById(ctx context.Context, userId int) (*core.User, error) {
	var output core.User
	query := "SELECT * FROM users.users WHERE id = $1"
	err := r.db.QueryRowx(query, userId).StructScan(&output)
	return &output, err
}

func (r *users) GetByToken(ctx context.Context, token string) (*core.User, error) {
	var output core.User
	query := "SELECT * FROM users.users WHERE id = (SELECT user_id FROM users.sessions WHERE token = $1)"
	err := r.db.QueryRowx(query, token).StructScan(&output)
	return &output, err
}

func (r *users) GetByEmail(ctx context.Context, email string) (*core.User, error) {
	var output core.User
	sql := "SELECT * FROM users.users WHERE email = $1"
	err := r.db.QueryRowx(sql, email).StructScan(&output)
	return &output, err
}

func (r *users) EmailExists(ctx context.Context, email string) error {
	var userId int
	query := "SELECT id FROM users.users WHERE email = $1"
	err := r.db.QueryRow(query, email).Scan(&userId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func (r *users) CreateToken(ctx context.Context, userId int) (string, error) {
	token, err := core.GenerateRandomString(128)
	if err != nil {
		return "", nil
	}

	query := "INSERT INTO users.sessions (user_id, token) VALUES ($1, $2);"
	err = r.db.QueryRow(query, userId, token).Scan()
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	return token, nil
}

func (r *users) Update(ctx context.Context, user *core.User) error {
	query := "UPDATE users.users SET email = $1, phone = $2, firstname = $3, lastname = $4, patronymic = $5 WHERE id = $6"
	err := r.db.QueryRow(query, user.Email, user.Phone, user.FirstName, user.LastName, user.Patronymic, user.ID).Scan()
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func (r *users) SendRecoveryKey(ctx context.Context, email string) error {
	user, err := r.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	key, err := core.GenerateRandomString(128)
	if err != nil {
		return err
	}

	expire := time.Now().Add(12 * time.Hour).UTC()

	query := "INSERT INTO users.recovery_keys (user_id, email, key, expire) VALUES ($1, $2, $3, $4);"
	if err := r.db.QueryRow(query, user.ID, email, key, expire).Scan(); err != nil && err != sql.ErrNoRows {
		return err
	}

	// mailContent := map[string]string{
	// 	"key": key,
	// }
	// if err := r.mailService.Send(ctx, email, "recovery", mailContent); err != nil {
	// 	return err
	// }

	return nil
}

func (r *users) DeleteRecoveryKey(ctx context.Context, key string) error {
	query := "DELETE FROM users.recovery_keys WHERE key = $1"
	err := r.db.QueryRow(query, key).Scan()
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (r *users) GetUserIdByRecoveryKey(ctx context.Context, key string) (int, error) {
	var userId int
	query := "SELECT user_id FROM users.recovery_keys WHERE key = $1"
	err := r.db.QueryRow(query, key).Scan(&userId)
	return userId, err
}

func (r *users) SetPasswordByUserId(ctx context.Context, userId int, password string) error {
	query := "UPDATE users.users SET password = $1 WHERE id = $2"
	err := r.db.QueryRow(query, password, userId).Scan()
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}
