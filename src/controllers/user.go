package controllers

import (
	"log"
	"net/http"

	"github.com/adrimm6661604086/TPV_Bank-Simulator/database"
	"github.com/adrimm6661604086/TPV_Bank-Simulator/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/** GetUserByID
* obtiene el usuario por su ID
*
* @param c: contexto de la petición
* @return user Struct: usuario
 */
func GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if database.DB == nil {
		log.Println("Database connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal server error",
		})
		return
	}

	result := database.DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Println("User not found")
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
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
		"data":    user,
	})
}
