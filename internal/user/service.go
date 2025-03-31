package user

import (
	"context"
	"github.com/google/uuid"
)

type InterfaceService interface {
	CreateUser(ctx context.Context, data CreateUserRequest) error
}

type Service struct {
	InterfaceRepository InterfaceRepository
}

func NewServiceUser(InterfaceRepository InterfaceRepository) *Service {
	return &Service{
		InterfaceRepository,
	}
}

func (s *Service) CreateUser(ctx context.Context, data CreateUserRequest) error {
	arg := data.ParseCreateToUser()

	generate, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	arg.ID = generate
	err = s.InterfaceRepository.CreateUser(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}
