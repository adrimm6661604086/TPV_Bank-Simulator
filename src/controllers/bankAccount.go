package controllers

import (
	"log"
	"net/http"

	"github.com/adrimm6661604086/TPV_Bank-Simulator/database"
	"github.com/adrimm6661604086/TPV_Bank-Simulator/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/** GetAccounts
* obtiene todas las cuentas bancarias
*
* @param c: contexto de la petición
* @return bankAccounts Struct: lista de cuentas bancarias
 */
func GetAccounts(c *gin.Context) {
	var bankAccounts []models.BankAccount

	if database.DB == nil {
		log.Println("Database connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal server error",
		})
		return
	}

	result := database.DB.Find(&bankAccounts)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Println("Bank Account not found")
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Bank Account not found",
				"data":    nil,
			})
			return
		}

		log.Println("Database error:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    bankAccounts,
	})
}

/** GetAccountByID
* obtiene una cuenta bancaria por su ID
*
* @param
* 		c: contexto de la petición
* 		id: ID de la cuenta bancaria a obtener
* @return bankAccount Struct: cuenta bancaria obtenida
 */
func GetAccountByID(c *gin.Context) {
	id := c.Param("id")

	var bankAccount models.BankAccount
	result := database.DB.First(&bankAccount, "id = ?", id).Select("ID", "UserID", "IBAN", "Balance")

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Println("Bank Account not found for ID:", id)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Bank Account not found",
				"data":    nil,
			})
			return
		}

		log.Println("Database error:", result.Error)
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Internal server error",
				"data":    nil,
			})
		return
	}

	c.JSON(http.StatusOK, bankAccount)
}
