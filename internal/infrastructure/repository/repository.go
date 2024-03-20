package repository

import (
	"api/internal/core"
	"api/internal/infrastructure/repository/users"
	"context"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	Users IUsers
}

type IUsers interface {
	Create(ctx context.Context, user *core.User) (int, error)
	GetOneById(ctx context.Context, userId int) (*core.User, error)
	GetByToken(ctx context.Context, token string) (*core.User, error)
	GetByEmail(ctx context.Context, email string) (*core.User, error)
	EmailExists(ctx context.Context, email string) error
	CreateToken(ctx context.Context, userId int) (string, error)
	Update(ctx context.Context, user *core.User) error
	UpdatePhoto(ctx context.Context, userId int, path string) error

	SendRecoveryKey(ctx context.Context, email string) error
	GetUserIdByRecoveryKey(ctx context.Context, key string) (int, error)
	DeleteRecoveryKey(ctx context.Context, key string) error
	SetPasswordByUserId(ctx context.Context, userId int, password string) error
}

func New(db *sqlx.DB) *Repo {
	return &Repo{
		Users: users.New(db),
	}
}
