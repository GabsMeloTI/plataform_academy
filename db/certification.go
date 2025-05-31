package db

import "github.com/google/uuid"

type Certification struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name string    `gorm:"type:varchar(255);unique;not null" json:"name"`
}
