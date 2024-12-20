package controllers

import (
	"errors"
	"log"

	"github.com/adrimm6661604086/TPV_Bank-Simulator/database"
)

/**
* Process the payment on the bank simulator
*
* @param IBANorig: IBAN of the origin account
* @param IBANdst: IBAN of the destination account
* @param amount: amount of money to transfer
* @return {}
 */
func ProcessPayment(IBANorig string, IBANdst string, amount float64) (map[string]string, error) {
	log.Println("Processing payment from", IBANorig, "to", IBANdst, "of", amount)

	// Validar entradas
	if IBANorig == "" || IBANdst == "" || amount <= 0 {
		log.Fatalln("Invalid payment data")
		return nil, errors.New("Invalid payment data")
	}

	if IBANorig == IBANdst {
		log.Fatalln("Origin and destination accounts are the same")
		return nil, errors.New("Origin and destination accounts cannot be the same")
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		log.Fatalln("Failed to begin transaction:", tx.Error)
		return nil, tx.Error
	}

	var originBalance float64
	if err := tx.Raw("SELECT balance FROM bank_accounts WHERE IBAN = ?", IBANorig).Scan(&originBalance).Error; err != nil {
		log.Fatalln("Failed to retrieve origin account balance:", err)
		tx.Rollback()
		return nil, errors.New("Origin account not found")
	}

	if originBalance < amount {
		log.Fatalln("Insufficient balance in origin account")
		tx.Rollback()
		return nil, errors.New("Insufficient balance in origin account")
	}

	if err := tx.Exec("UPDATE accounts SET balance = balance - ? WHERE IBAN = ?", amount, IBANorig).Error; err != nil {
		log.Fatalln("Failed to debit origin account:", err)
		tx.Rollback()
		return nil, errors.New("Failed to debit origin account")
	}

	if err := tx.Exec("UPDATE accounts SET balance = balance + ? WHERE IBAN = ?", amount, IBANdst).Error; err != nil {
		log.Fatalln("Failed to credit destination account:", err)
		tx.Rollback()
		return nil, errors.New("Failed to credit destination account")
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed to commit transaction:", err)
		return nil, errors.New("Failed to process payment")
	}

	log.Println("Payment processed successfully")
	return map[string]string{
		"message": "Payment processed successfully",
		"status":  "OK",
	}, nil
}
