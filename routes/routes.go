package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the CRM system!",
		})
	})
	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "List of users",
		})
	})
	router.GET("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "User details",
		})
	})
}
