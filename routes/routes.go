package routes

import (
	controller "crm/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userController *controller.UserController) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the CRM system!",
		})
	})

	router.GET("/users", userController.GetAllUsers)
	router.GET("/users/:id", userController.GetUser)
	router.POST("/users", userController.CreateUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)
}
