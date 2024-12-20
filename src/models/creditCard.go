package models

import (
	"github.com/google/uuid"
)

const (
	Visa            = "Visa"
	Mastercard      = "Mastercard"
	AmericanExpress = "American Express"
)

type CardSchema string

type CreditCard struct {
	ID               uuid.UUID  `gorm:"type:uuid;primary_key;"`
	IBAN             string     `gorm:"type:varchar(34);not null;"`
	CardSchema       CardSchema `gorm:"type:varchar(20);not null;"`
	CreditCardNumber string     `gorm:"type:varchar(16);not null;"`
	PIN              string     `gorm:"type:varchar(4);not null;"`
	CVC              string     `gorm:"type:varchar(3);not null;"`
	Cardholder       string     `gorm:"type:varchar(100);not null;"`
	ExpirationDate   string     `gorm:"type:date;not null;"`
}

func (CreditCard) TableName() string {
	return "credit_cards"
}
