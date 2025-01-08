package controllers

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"

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

	if IBANorig == "" || IBANdst == "" || amount <= 0 {
		log.Println("Invalid payment data")
		return nil, errors.New("invalid payment data")
	}

	if IBANorig == IBANdst {
		log.Println("Origin and destination accounts are the same")
		return nil, errors.New("origin and destination accounts cannot be the same")
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		log.Println("Failed to begin transaction:", tx.Error)
		return nil, tx.Error
	}

	var originBalance float64
	row := tx.Raw("SELECT balance FROM bank_accounts WHERE REPLACE(iban, ' ', '') = REPLACE(?, ' ', '')", IBANorig).Row()
	if err := row.Scan(&originBalance); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("Origin account not found for IBAN:", IBANorig)
			tx.Rollback()
			return nil, errors.New("origin account not found")
		}
		log.Println("Error querying origin account balance:", err)
		tx.Rollback()
		return nil, errors.New("failed to retrieve origin account balance")
	}

	log.Println("Origin account balance:", originBalance)
	if originBalance < amount {
		log.Println("Insufficient balance in origin account")
		tx.Rollback()
		return nil, errors.New("insufficient balance in origin account")
	}

	if err := tx.Exec("UPDATE bank_accounts SET balance = balance + ? WHERE REPLACE(iban, ' ', '') = REPLACE(?, ' ', '')", amount, IBANorig).Error; err != nil {
		log.Println("Failed to debit origin account:", err)
		tx.Rollback()
		return nil, errors.New("failed to debit origin account")
	}

	if err := tx.Exec("UPDATE bank_accounts SET balance = balance - ? WHERE REPLACE(iban, ' ', '') = REPLACE(?, ' ', '')", amount, IBANdst).Error; err != nil {
		log.Println("Failed to credit destination account:", err)
		tx.Rollback()
		return nil, errors.New("failed to credit destination account")
	}

	transactionId := uuid.New().String()
	if err := tx.Exec("INSERT INTO transactions (id, origin_account, destination_account, amount, date) VALUES (?, ?, ?, ?, ?)",
		transactionId, IBANorig, IBANdst, amount, time.Now()).Error; err != nil {
		log.Println("Failed to register transaction:", err)
		tx.Rollback()
		return nil, errors.New("failed to register transaction")
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed to commit transaction:", err)
		return nil, errors.New("failed to process payment")
	}

	log.Println("Payment processed successfully")
	return map[string]string{
		"message": "Payment processed successfully",
		"status":  "OK",
	}, nil
}
