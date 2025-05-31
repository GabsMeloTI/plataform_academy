package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Trainer struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID             uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Cref               string    `gorm:"type:varchar(50)" json:"cref,omitempty"`
	College            string    `gorm:"type:varchar(255)" json:"college,omitempty"`
	GraduationDate     time.Time `gorm:"type:date" json:"graduation_date,omitempty"`
	ExperienceDuration string    `gorm:"type:varchar(50)" json:"experience_duration,omitempty"`
	Specialty          string    `gorm:"type:varchar(255)" json:"specialty,omitempty"`
	Certifications     string    `gorm:"type:varchar(255)" json:"certifications,omitempty"`
	PrimarySpecialty   string    `gorm:"type:varchar(255)" json:"primary_specialty,omitempty"`
	Bio                string    `gorm:"type:text" json:"bio,omitempty"`
	gorm.Model
}

type CreateTrainer struct {
	UserID             uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Cref               string    `gorm:"type:varchar(50)" json:"cref,omitempty"`
	College            string    `gorm:"type:varchar(255)" json:"college,omitempty"`
	GraduationDate     time.Time `gorm:"type:date" json:"graduation_date,omitempty"`
	ExperienceDuration string    `gorm:"type:varchar(50)" json:"experience_duration,omitempty"`
	Specialty          string    `gorm:"type:varchar(255)" json:"specialty,omitempty"`
	Certifications     string    `gorm:"type:varchar(255)" json:"certifications,omitempty"`
	PrimarySpecialty   string    `gorm:"type:varchar(255)" json:"primary_specialty,omitempty"`
	Bio                string    `gorm:"type:text" json:"bio,omitempty"`
}

type UpdateTrainer struct {
	UserID             uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Cref               string    `gorm:"type:varchar(50)" json:"cref,omitempty"`
	College            string    `gorm:"type:varchar(255)" json:"college,omitempty"`
	GraduationDate     time.Time `gorm:"type:date" json:"graduation_date,omitempty"`
	ExperienceDuration string    `gorm:"type:varchar(50)" json:"experience_duration,omitempty"`
	Specialty          string    `gorm:"type:varchar(255)" json:"specialty,omitempty"`
	Certifications     string    `gorm:"type:varchar(255)" json:"certifications,omitempty"`
	PrimarySpecialty   string    `gorm:"type:varchar(255)" json:"primary_specialty,omitempty"`
	Bio                string    `gorm:"type:text" json:"bio,omitempty"`
}

func (t *CreateTrainer) ParseCreateToTrainer() Trainer {
	arg := Trainer{
		UserID:             t.UserID,
		Cref:               t.Cref,
		College:            t.College,
		GraduationDate:     t.GraduationDate,
		ExperienceDuration: t.ExperienceDuration,
		Specialty:          t.Specialty,
		Certifications:     t.Certifications,
		PrimarySpecialty:   t.PrimarySpecialty,
		Bio:                t.Bio,
		Model:              gorm.Model{},
	}

	return arg
}
