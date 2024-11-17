package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConnection string
	DATABASE_URL string
	ServerPort   string
}

type ConfigDB struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	SSLMode    string
}

func LoadConfig() Config {
	log.SetPrefix("Setting Config(): ")
	log.Println("Loading environment variables ...")
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	var configDB ConfigDB = ConfigDB{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		SSLMode:    os.Getenv("SSL_MODE"),
	}

	return Config{
		DBConnection: configDB.GetDBConnectionString(),
		DATABASE_URL: os.Getenv("DATABASE_URL"),
		ServerPort:   os.Getenv("TPV_BANK_PORT"),
	}
}

func (c ConfigDB) GetDBConnectionString() string {
	return "host=" + c.DBHost +
		" user=" + c.DBUser +
		" password='" + c.DBPassword + "'" +
		" dbname=" + c.DBName +
		" port=" + c.DBPort +
		" sslmode=" + c.SSLMode
}
