package database

import (
	"log"

	// "gorm.io/gorm/logger"
	"github.com/adrimm6661604086/TPV_Bank-Simulator/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(connectionString string) {
	originalPrefix := log.Prefix()
	log.SetPrefix("Dataabase connection: ")
	defer log.SetPrefix(originalPrefix)

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	// Migración de modelos a la base de datos para tener las tablas up-to-date
	if err := DB.AutoMigrate(&models.User{}, &models.BankAccount{}, &models.CreditCard{}, &models.Transaction{}); err != nil {
		log.Fatal("Error migrating models: ", err)
	}

	log.Println("Connected to database")
}
