package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"plataform_init/db"
)

type InterfaceService interface {
	CreateUser(ctx context.Context, data db.CreateUser) error
	Login(ctx context.Context, data db.LoginUser) (string, error)
}

type Service struct {
	InterfaceRepository InterfaceRepository
}

func NewServiceUser(InterfaceRepository InterfaceRepository) *Service {
	return &Service{
		InterfaceRepository,
	}
}

func (s *Service) CreateUser(ctx context.Context, data db.CreateUser) error {
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

func (s *Service) Login(ctx context.Context, data db.LoginUser) (string, error) {
	user, err := s.InterfaceRepository.FindUserByEmail(ctx, data.Email)
	if err != nil {
		return "", err
	}

	if user.ID == uuid.Nil {
		return "", errors.New("usuário não encontrado")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return "", errors.New("senha incorreta")
	}

	return "token_de_exemplo", nil
}
