package main

import (
	"flag"
	"log"
	"os"

	"github.com/adrimm6661604086/TPV_Bank-Simulator/config"
	"github.com/adrimm6661604086/TPV_Bank-Simulator/database"
	"github.com/adrimm6661604086/TPV_Bank-Simulator/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Definir la flag para el entorno
	env := flag.String("env", "dev", "Define the environment: dev or prod")
	flag.Parse()

	if *env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	logger, err := os.OpenFile("logger.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logger)
	defer logger.Close()

	var config config.Config = config.LoadConfig()

	database.ConnectDB(config.DBConnection)

	var router *gin.Engine = routes.Router(config)

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal(err)
	}

	log.SetPrefix("INFO: ")
	log.Println("Server running on http://localhost:", config.ServerPort)
	if err := router.Run("127.0.0.1:" + config.ServerPort); err != nil {
		log.Fatal("Server error: ", err)
	}
}
