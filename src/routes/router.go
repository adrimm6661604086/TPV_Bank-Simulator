package routes

import (
	"os"
	"path/filepath"

	"github.com/adrimm6661604086/TPV_Bank-Simulator/config"
	"github.com/gin-gonic/gin"
)

func Router(config config.Config) *gin.Engine {
	var router *gin.Engine = gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the TPV Bank Simulator running on http://localhost:" + config.ServerPort})
	})

	router.GET("/logger", func(c *gin.Context) {
		wd, err := os.Getwd()
		if err != nil {
			c.String(500, "Error getting current directory")
			return
		}

		logFilePath := filepath.Join(wd, "logger.log")
		c.File(logFilePath)
	})

	BankAccountRoutes(router)
	CreditCardRoutes(router)
	TransactionRoutes(router)
	UserRoutes(router)

	return router
}
