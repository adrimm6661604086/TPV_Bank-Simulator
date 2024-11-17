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
func GetCreditCards(c *gin.Context) {
	var creditCards []models.CreditCard

	if database.DB == nil {
		log.Println("Database connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal server error",
		})
		return
	}

	result := database.DB.Find(&creditCards)
	if result.Error != nil {
		log.Println("Database error:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal server error",
		})
		return
	}

	if result.RowsAffected == 0 {
		log.Println("BankAccount Controller: No bank accounts found")
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No bank accounts found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    creditCards,
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
func GetCreditCardByNumber(c *gin.Context) {
	id := c.Param("id")

	// Aquí implementas la lógica para buscar la cuenta en la base de datos
	// Este es un ejemplo de respuesta simulada:
	var creditCard models.CreditCard
	result := database.DB.First(&creditCard, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Println("Bank Account not found for ID:", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "Bank Account not found"})
			return
		}

		log.Println("Database error:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, creditCard)
}
