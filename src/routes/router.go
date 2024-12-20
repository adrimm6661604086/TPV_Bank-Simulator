package routes

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/adrimm6661604086/TPV_Bank-Simulator/config"
	"github.com/adrimm6661604086/TPV_Bank-Simulator/src/controllers"
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

	router.POST("/process-payment", func(c *gin.Context) {
		creditCardNumber, _ := c.GetPostForm("creditCardNumber")
		PIN, _ := c.GetPostForm("PIN")
		creditCardHolder, _ := c.GetPostForm("creditCardHolder")
		expirationDate, _ := c.GetPostForm("expirationDate")
		CVC, _ := c.GetPostForm("CVC")

		if creditCardNumber == "" || PIN == "" || creditCardHolder == "" || expirationDate == "" || CVC == "" {
			c.JSON(400, gin.H{"message": "Missing credit card credentials"})
			return
		}

		checkCard := controllers.VerifyCreditCard(creditCardNumber, PIN, creditCardHolder, expirationDate, CVC)

		if checkCard.IsValid {
			IBANorig, _ := c.GetPostForm("IBANorig")
			amountStr, _ := c.GetPostForm("amount")
			amount, err := strconv.ParseFloat(amountStr, 64)

			if err != nil {
				c.JSON(400, gin.H{"message": "Invalid amount"})
				return
			}

			controllers.ProcessPayment(IBANorig, checkCard.IBANdst, amount)

			c.JSON(200, gin.H{"message": "Credit card credentials verified"})
		} else {
			c.JSON(400, gin.H{"message": "Credit card credentials not verified"})
		}
	})

	return router
}
