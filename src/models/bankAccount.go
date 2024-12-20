package models

import (
	"github.com/google/uuid"
)

type BankAccount struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID  uuid.UUID `gorm:"type:uuid;not null;"`
	IBAN    string    `gorm:"type:varchar(34);unique;not null;"`
	Balance float64   `gorm:"type:float;not null;check:balance >= 0;"`
	User    User      `gorm:"foreignKey:UserID;references:ID; json:"-"` //constraint:OnUpdate:CASCADE,OnDelete:CASCADE;
}

func (BankAccount) TableName() string {
	return "bank_accounts"
}
