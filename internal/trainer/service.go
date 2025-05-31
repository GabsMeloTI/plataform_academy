package trainer

import (
	"context"
	"github.com/google/uuid"
	"plataform_init/db"
)

type InterfaceService interface {
	CreateUser(ctx context.Context, data db.CreateTrainer) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) CreateUser(ctx context.Context, data db.CreateTrainer, userId uuid.UUID) error {
	arg := data.ParseCreateToTrainer()
	arg.UserID = userId

	err := s.repo.CreateTrainer(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}
