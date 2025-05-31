package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"path"
	"plataform_init/infra/token"
	"plataform_init/pkg"
	"plataform_init/validation"
	"strings"
	"time"

	"github.com/google/uuid"
	"plataform_init/db"
)

type InterfaceService interface {
	CreateUser(ctx context.Context, data db.CreateUser) error
	Login(ctx context.Context, data db.LoginUser) (string, error)
	UpdateAvatar(ctx context.Context, data *multipart.Form, id uuid.UUID) (string, error)
	GetUserById(ctx context.Context, id uuid.UUID) (db.UserList, error)
	GetUsersByRole(ctx context.Context, role string) (interface{}, error)
}

type Service struct {
	InterfaceRepository InterfaceRepository
	secretKey           string
	bucket              string
}

func NewServiceUser(InterfaceRepository InterfaceRepository, secretKey string, bucket string) *Service {
	return &Service{
		InterfaceRepository,
		secretKey,
		bucket,
	}
}

func (s *Service) CreateUser(ctx context.Context, data db.CreateUser) error {
	if err := validation.ValidateEmail(data.Email); err != nil {
		return err
	}

	existingUser, err := s.InterfaceRepository.FindUserByEmail(ctx, data.Email)
	if err == nil && existingUser.ID != uuid.Nil {
		return errors.New("e-mail já cadastrado")
	}

	ok := validation.ValidatePassword(data.Password)
	if !ok {
		return errors.New("senha invalida")
	}

	arg := data.ParseCreateToUser()

	arg.ID, err = GenerateUUID()
	if err != nil {
		return err
	}

	arg.Password, err = HashPassword(data.Password)
	if err != nil {
		return err
	}

	err = s.InterfaceRepository.CreateUser(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Login(ctx context.Context, data db.LoginUser) (string, error) {
	user, err := s.InterfaceRepository.FindUserByEmail(ctx, data.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("usuário não encontrado")
		}
		return "", err
	}

	if !CheckPasswordHash(data.Password, user.Password) {
		return "", errors.New("credênciais inválidas")
	}

	jwtMaker, err := token.NewJwtMaker(s.secretKey)
	if err != nil {
		log.Fatal("Erro ao criar o maker de JWT:", err)
	}

	result, err := jwtMaker.CreateToken(user.ID, user.Username, user.Email, time.Now().Add(24*time.Hour).UTC())
	if err != nil {
		log.Fatal("Erro ao criar o token:", err)
	}

	return result, nil
}

func (s *Service) UpdateAvatar(ctx context.Context, data *multipart.Form, id uuid.UUID) (string, error) {
	var myForm db.UpdateAvatar
	err := MapFormToStruct(data.Value, &myForm)
	if err != nil {
		return "", err
	}

	files, ok := data.File["avatar"]
	if !ok || len(files) == 0 {
		return "", fmt.Errorf("arquivo de avatar não encontrado")
	}

	fileHeader := files[0]
	idFile := GetUUID()
	originalFilename := fileHeader.Filename
	fileExtension := strings.ToLower(path.Ext(originalFilename))
	newNameFileUp := idFile + fileExtension

	f, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	contentType := fileHeader.Header.Get("Content-Type")

	upload, err := pkg.UploadFileToS3(fileBytes, newNameFileUp, s.bucket, contentType)
	if err != nil {
		return "", err
	}

	myForm.Avatar = upload
	err = s.InterfaceRepository.UpdateAvatar(ctx, myForm, id)
	if err != nil {
		return "", err
	}

	return upload, nil
}

func (s *Service) GetUserById(ctx context.Context, id uuid.UUID) (db.UserList, error) {
	result, err := s.InterfaceRepository.GetUserByID(ctx, id)
	if err != nil {
		return db.UserList{}, err
	}

	user := db.UserList{}
	user.ParseUserListToUser(result)

	return user, nil
}

func (s *Service) GetUsersByRole(ctx context.Context, role string) (interface{}, error) {
	if role != "student" && role != "personal" {
		return nil, errors.New("invalid role, it can only be 'student' or 'personal'")
	}

	result, err := s.InterfaceRepository.GetUsersByRole(ctx, role)
	if err != nil {
		return nil, err
	}

	if role == "student" {
		var students []db.UserList
		for _, u := range result {
			user := db.UserList{
				ID:          u.UserID,
				Name:        u.Name,
				Username:    u.Username,
				Email:       u.Email,
				Phone:       u.Phone,
				DateOfBirth: u.DateOfBirth,
				Gender:      u.Gender,
				Role:        u.Role,
				Status:      u.Status,
				Weight:      u.Weight,
				Height:      u.Height,
				Avatar:      u.Avatar,
			}
			students = append(students, user)
		}
		return students, nil
	}

	var personals []db.UsersCompleteResponse
	for _, u := range result {
		userPersonal := db.UsersCompleteResponse{
			UserID:             u.UserID,
			Name:               u.Name,
			Username:           u.Username,
			Email:              u.Email,
			Phone:              u.Phone,
			DateOfBirth:        u.DateOfBirth,
			Gender:             u.Gender,
			Role:               u.Role,
			Status:             u.Status,
			Weight:             u.Weight,
			Height:             u.Height,
			Avatar:             u.Avatar,
			Cref:               u.Cref,
			TrainerID:          u.TrainerID,
			College:            u.College,
			GraduationDate:     u.GraduationDate,
			ExperienceDuration: u.ExperienceDuration,
			Specialty:          u.Specialty,
			Certifications:     u.Certifications,
			PrimarySpecialty:   u.PrimarySpecialty,
			Bio:                u.Bio,
		}

		personals = append(personals, userPersonal)
	}

	return personals, nil
}
