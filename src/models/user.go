package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name           string    `gorm:"type:varchar(20);not null;"`
	Lastname       string    `gorm:"type:varchar(20);not null;"`
	SecondLastname string    `gorm:"type:varchar(20);not null;"`
	Password       string    `gorm:"type:varchar(128);not null;"`
	DNI            string    `gorm:"type:varchar(10);unique;not null;"`
	Country        string    `gorm:"type:varchar(20);not null;"`
}

func (User) TableName() string {
	return "users"
}
