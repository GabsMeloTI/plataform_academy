package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Progress struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ClientID       uuid.UUID `gorm:"type:uuid;not null" json:"client_id"`
	TrainerID      uuid.UUID `gorm:"type:uuid;not null" json:"trainer_id"`
	RecordDate     time.Time `gorm:"type:date;not null" json:"record_date"`
	Weight         float64   `gorm:"type:decimal(5,2)" json:"weight,omitempty"`
	BodyFatPercent float64   `gorm:"type:decimal(5,2)" json:"body_fat_percent,omitempty"`
	MusclePercent  float64   `gorm:"type:decimal(5,2)" json:"muscle_percent,omitempty"`
	Notes          string    `gorm:"type:text" json:"notes,omitempty"`
	gorm.Model
}

type CreateProgress struct {
	ClientID       uuid.UUID `gorm:"type:uuid;not null" json:"client_id"`
	TrainerID      uuid.UUID `gorm:"type:uuid;not null" json:"trainer_id"`
	RecordDate     time.Time `gorm:"type:date;not null" json:"record_date"`
	Weight         float64   `gorm:"type:decimal(5,2)" json:"weight,omitempty"`
	BodyFatPercent float64   `gorm:"type:decimal(5,2)" json:"body_fat_percent,omitempty"`
	MusclePercent  float64   `gorm:"type:decimal(5,2)" json:"muscle_percent,omitempty"`
	Notes          string    `gorm:"type:text" json:"notes,omitempty"`
}
