package trainer

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"plataform_init/db"
)

type InterfaceRepository interface {
	CreateTrainer(ctx context.Context, Trainer db.Trainer) error
	UpdateTrainer(ctx context.Context, Trainer db.UpdateTrainer) error
	GetTrainerByID(ctx context.Context, id uuid.UUID) (db.Trainer, error)
}

type Repository struct {
	dbGorm *gorm.DB
}

func NewTrainerRepository(dbGorm *gorm.DB) *Repository {
	return &Repository{
		dbGorm,
	}
}

func (r *Repository) CreateTrainer(ctx context.Context, Trainer db.Trainer) error {
	result := r.dbGorm.WithContext(ctx).Create(&Trainer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) UpdateTrainer(ctx context.Context, Trainer db.UpdateTrainer) error {
	result := r.dbGorm.WithContext(ctx).Model(&db.Trainer{}).Where("user_id = ?", Trainer.UserID).Updates(Trainer)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) GetTrainerByID(ctx context.Context, id uuid.UUID) (db.Trainer, error) {
	var Trainer db.Trainer
	result := r.dbGorm.WithContext(ctx).Where("id = ?", id).First(&Trainer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Trainer, errors.New("Trainer not found")
		}
		return Trainer, result.Error
	}

	return Trainer, nil
}
