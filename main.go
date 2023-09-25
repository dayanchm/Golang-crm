package main

import (
	"crm/controller"
	"crm/database"
	"crm/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// Veritabanı bağlantısını açın
	db, err := database.OpenDB()
	if err != nil {
		panic(err)
	}

	// Veritabanı bağlantısını kapattığınızdan emin olun
	defer db.Close()

	// UserController kontrolcüsünü oluşturun
	userController := controller.NewUserControllerWithDB(db)

	// Gin rota motorunu oluşturun
	r := gin.Default()

	// Rotaları ayarlayın
	routes.SetupRoutes(r, userController)

	// Uygulamayı belirli bir bağlantı noktasında başlatın
	err = r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
