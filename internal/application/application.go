package application

import (
	"api/internal/application/estates"
	"api/internal/application/users"
	"api/internal/core"
	"api/internal/infrastructure/repository"
	"context"
)

type UseCase struct {
	Users   IUsers
	Estates IEstates
}

type IUsers interface {
	SignUp(ctx context.Context, email, password string) (string, error)
	SignIn(ctx context.Context, email, password string) (string, error)

	SendRecoveryKey(ctx context.Context, email string) error
	ConfirmRecoveryKey(ctx context.Context, key, password string) (string, error)

	GetOneById(ctx context.Context, userId int) (*core.User, error)
	GetByToken(ctx context.Context, token string) (*core.User, error)
	Update(ctx context.Context, user *core.User) error
	UpdatePhoto(ctx context.Context, userId int, bytes []byte) (image string, err error)
}

type IBooking interface {
	Add(ctx context.Context, estate *core.Estate) error
	UpdateStatus(ctx context.Context, estateId int) error
}

type IEstates interface {
	Add(ctx context.Context, estate *core.Estate) error
	GetAll(ctx context.Context) (*[]core.Estate, error)
	GetById(ctx context.Context, id int) (*core.Estate, error)
	// Remove(ctx context.Context, id int) error
}

type IMessages interface {
	Add(ctx context.Context) error
}

func New(repo *repository.Repo) *UseCase {
	return &UseCase{
		Users:   users.New(repo.Users),
		Estates: estates.New(repo.Estates),
	}
}
