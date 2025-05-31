package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Appointment struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	TrainerID uuid.UUID `gorm:"type:uuid;not null" json:"trainer_id"`
	ClientID  uuid.UUID `gorm:"type:uuid;not null" json:"client_id"`
	StartTime time.Time `gorm:"type:timestamptz;not null" json:"start_time"`
	EndTime   time.Time `gorm:"type:timestamptz;not null" json:"end_time"`
	Status    string    `gorm:"type:varchar(20);not null;default:'scheduled'" json:"status"`
	Location  string    `gorm:"type:varchar(255)" json:"location,omitempty"`
	Notes     string    `gorm:"type:text" json:"notes,omitempty"`
	gorm.Model
}

type CreateAppointment struct {
	TrainerID uuid.UUID `gorm:"type:uuid;not null" json:"trainer_id"`
	ClientID  uuid.UUID `gorm:"type:uuid;not null" json:"client_id"`
	StartTime time.Time `gorm:"type:timestamptz;not null" json:"start_time"`
	EndTime   time.Time `gorm:"type:timestamptz;not null" json:"end_time"`
	Status    string    `gorm:"type:varchar(20)" json:"status,omitempty"`
	Location  string    `gorm:"type:varchar(255)" json:"location,omitempty"`
	Notes     string    `gorm:"type:text" json:"notes,omitempty"`
}
