package user

import (
	"context"
	"errors"
	"regexp"

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

func validateEmail(email string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("e-mail inválido")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("a senha deve ter pelo menos 6 caracteres")
	}
	return nil
}

func (s *Service) CreateUser(ctx context.Context, data db.CreateUser) error {
	if err := validateEmail(data.Email); err != nil {
		return err
	}

	existingUser, err := s.InterfaceRepository.FindUserByEmail(ctx, data.Email)
	if err == nil && existingUser.ID != uuid.Nil {
		return errors.New("e-mail já cadastrado")
	}

	if err := validatePassword(data.Password); err != nil {
		return err
	}

	generate, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	arg := data.ParseCreateToUser()
	arg.ID = generate

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	arg.Password = string(hashedPassword)

	err = s.InterfaceRepository.CreateUser(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Login(ctx context.Context, data db.LoginUser) (string, error) {
	if err := validateEmail(data.Email); err != nil {
		return "", err
	}

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
