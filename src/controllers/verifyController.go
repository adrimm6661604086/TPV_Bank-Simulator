package controllers

import (
	"errors"
	"log"

	"github.com/adrimm6661604086/TPV_Bank-Simulator/database"
	"github.com/adrimm6661604086/TPV_Bank-Simulator/src/models"
	"gorm.io/gorm"
)

type CreditCardVerificationResult struct {
	IBANdst string
	IsValid bool
	Error   string
}

/**
* VerifyCreditCard
* Verifica la información de la tarjeta de crédito
*
* @param creditCardNumber: número de tarjeta de crédito
* @param PIN: PIN de la tarjeta de crédito
* @param creditCardHolder: titular de la tarjeta de crédito
* @param expirationDate: fecha de expiración de la tarjeta de crédito
* @param CVC: CVC de la tarjeta de crédito
* @return
 */
func VerifyCreditCard(creditCardNumber string, PIN string, creditCardHolder string, expirationDate string, CVC string) CreditCardVerificationResult {
	log.Println("Verifying credit card credentials")

	if len(creditCardNumber) == 0 || len(PIN) == 0 || len(creditCardHolder) == 0 || len(expirationDate) == 0 || len(CVC) == 0 {
		log.Fatalln("Invalid input parameters")
		return CreditCardVerificationResult{
			IBANdst: "",
			IsValid: false,
			Error:   "Invalid input parameters",
		}
	}

	var cCard models.CreditCard

	if err := database.DB.Where("credit_card_number = ?", creditCardNumber).First(&cCard).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Fatalln("Credit card not found")
			return CreditCardVerificationResult{
				IBANdst: "",
				IsValid: false,
				Error:   "Credit card not found",
			}
		}
		log.Fatalln("Database error:", err)
		return CreditCardVerificationResult{
			IBANdst: "",
			IsValid: false,
			Error:   "Database error",
		}
	}

	if cCard.PIN != PIN {
		log.Fatalln("Invalid PIN")
		return CreditCardVerificationResult{
			IBANdst: "",
			IsValid: false,
			Error:   "Invalid PIN",
		}
	}

	if cCard.CVC != CVC {
		log.Fatalln("Invalid CVC")
		return CreditCardVerificationResult{
			IBANdst: "",
			IsValid: false,
			Error:   "Invalid CVC",
		}
	}

	log.Println("Credit card verified successfully")
	return CreditCardVerificationResult{
		IBANdst: cCard.IBAN,
		IsValid: true,
		Error:   "",
	}
}
