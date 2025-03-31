package user

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"plataform_init/db"
)

type InterfaceRepository interface {
	CreateUser(ctx context.Context, user db.User) error
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
	fmt.Println(result)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
