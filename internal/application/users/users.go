package users

import (
	"api/internal/core"
	"api/internal/infrastructure/repository"
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type users struct {
	repo repository.IUsers
}

func New(repo repository.IUsers) *users {
	return &users{repo: repo}
}

func (s *users) SignUp(ctx context.Context, email, password string) (string, error) {
	err := s.repo.EmailExists(ctx, email)
	if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	user := &core.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	userId, err := s.repo.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return s.repo.CreateToken(ctx, userId)
}

func (s *users) SignIn(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	return s.repo.CreateToken(ctx, user.ID)
}

func (s *users) GetOneById(ctx context.Context, userId int) (*core.User, error) {
	return s.repo.GetOneById(ctx, userId)
}

func (s *users) GetByToken(ctx context.Context, token string) (*core.User, error) {
	return s.repo.GetByToken(ctx, token)
}

func (s *users) Update(ctx context.Context, user *core.User) error {
	return s.repo.Update(ctx, user)
}

func (s *users) SendRecoveryKey(ctx context.Context, email string) error {
	if err := s.repo.EmailExists(ctx, email); err != nil {
		return err
	}
	return s.repo.SendRecoveryKey(ctx, email)
}

func (s *users) ConfirmRecoveryKey(ctx context.Context, key, password string) (string, error) {
	userId, err := s.repo.GetUserIdByRecoveryKey(ctx, key)
	if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	if err := s.repo.SetPasswordByUserId(ctx, userId, string(hashedPassword)); err != nil {
		return "", err
	}

	if err := s.repo.DeleteRecoveryKey(ctx, key); err != nil {
		log.Println(err)
		return "", err
	}

	return s.repo.CreateToken(ctx, userId)
}
