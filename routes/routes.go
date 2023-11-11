package routes

import (
	controller "crm/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userController *controller.UserController) {
	router.LoadHTMLGlob("views/*")
	router.Static("assets", "./assets")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", gin.H{
			"title": "Main website",
		})
	})
	router.POST("/users", userController.CreateUser)
	router.DELETE("/users/:id", userController.DeleteUser)
	router.GET("/users/:id", userController.GetOneUser)
	router.GET("/users", userController.GetAllUsers)
	router.POST("/users/:id", userController.UpdateUser)

}
