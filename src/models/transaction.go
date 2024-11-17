package models

import (
	"github.com/google/uuid"
)

type Transaction struct {
	ID                 uuid.UUID `gorm:"type:uuid;primary_key;"`
	OriginAccount      string    `gorm:"type:varchar(34);not null;"`
	DestinationAccount string    `gorm:"type:varchar(34);not null;"`
	Amount             float64   `gorm:"type:float;not null;check:amount > 0;"`
	Date               string    `gorm:"type:date;not null;"`
}

func (Transaction) TableName() string {
	return "transactions"
}
