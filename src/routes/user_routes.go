package routes

import (
	"github.com/adrimm6661604086/TPV_Bank-Simulator/src/controllers"
	"github.com/gin-gonic/gin"
)

// UserRoutes define las rutas relacionadas con usuarios
func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		// userGroup.GET("/", controllers.GetAllUsers)
		userGroup.GET("/:id", controllers.GetUserByID)
	}
}
