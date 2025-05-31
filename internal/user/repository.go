package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"plataform_init/db"
)

type InterfaceRepository interface {
	CreateUser(ctx context.Context, user db.User) error
	FindUserByEmail(ctx context.Context, email string) (db.User, error)
	UpdateAvatar(ctx context.Context, user db.UpdateAvatar, id uuid.UUID) error
	GetUserByID(ctx context.Context, id uuid.UUID) (db.User, error)
	GetUsersByRole(ctx context.Context, role string) ([]db.UsersCompleteResponse, error)
}

type Repository struct {
	dbGorm *gorm.DB
}

func NewUserRepository(dbGorm *gorm.DB) *Repository {
	return &Repository{
		dbGorm,
	}
}

func (r *Repository) CreateUser(ctx context.Context, user db.User) error {
	result := r.dbGorm.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) FindUserByEmail(ctx context.Context, email string) (db.User, error) {
	var user db.User
	result := r.dbGorm.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, nil
		}
		return user, result.Error
	}
	return user, nil
}

func (r *Repository) UpdateAvatar(ctx context.Context, user db.UpdateAvatar, id uuid.UUID) error {
	result := r.dbGorm.WithContext(ctx).Model(&db.User{}).Where("id = ?", id).Update("avatar", user.Avatar)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) GetUserByID(ctx context.Context, id uuid.UUID) (db.User, error) {
	var user db.User
	result := r.dbGorm.WithContext(ctx).Where("id = ?", id).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
		return user, result.Error
	}

	return user, nil
}

func (r *Repository) GetUsersByRole(ctx context.Context, role string) ([]db.UsersCompleteResponse, error) {
	var usersPersonal []db.UsersCompleteResponse

	columns := `
		u.id AS user_id,
		u.name,
		u.username,
		u.email,
		u.password,
		u.phone,
		u.date_of_birth,
		u.gender,
		u.role,
		u.status,
		u.cref,
		u.avatar,
		t.id AS trainer_id,
		t.college,
		t.graduation_date,
		t.experience_duration,
		t.specialty,
		t.certifications,
		t.primary_specialty,
		t.bio
	`
	result := r.dbGorm.WithContext(ctx).
		Table("users AS u").
		Select(columns).
		Joins("INNER JOIN trainers AS t ON t.user_id = u.id").
		Where("u.role = ?", role).
		Find(&usersPersonal)

	if result.Error != nil {
		return nil, result.Error
	}

	return usersPersonal, nil
}
