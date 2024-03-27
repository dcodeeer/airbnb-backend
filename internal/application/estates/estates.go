package estates

import (
	"api/internal/core"
	"api/internal/infrastructure/repository"
	"context"
)

type estates struct {
	repo repository.IEstates
}

func New(repo repository.IEstates) *estates {
	return &estates{repo: repo}
}

func (s *estates) Add(ctx context.Context, estate *core.Estate) error {
	// тут должна быть транзакция
	return nil
}

func (s *estates) GetAll(ctx context.Context) (*[]core.Estate, error) {
	return s.repo.GetAll(ctx)
}

func (s *estates) GetById(ctx context.Context, id int) (*core.Estate, error) {
	return s.repo.GetOne(ctx, id)
}
