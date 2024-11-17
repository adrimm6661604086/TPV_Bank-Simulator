package routes

import (
	"github.com/adrimm6661604086/TPV_Bank-Simulator/src/controllers"
	"github.com/gin-gonic/gin"
)

func BankAccountRoutes(router *gin.Engine) {
	bankAccountGroup := router.Group("/bank-accounts")
	{
		bankAccountGroup.GET("/", controllers.GetAccounts)
		bankAccountGroup.GET("/:id", controllers.GetAccountByID)
	}
}
