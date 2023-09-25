package main

import (
	"crm/controller"
	"crm/database"
	"crm/model"
	"crm/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Gin'in çalışma modunu ayarla
	gin.SetMode(gin.ReleaseMode)

	// SQL veritabanı bağlantısını aç
	sqlDB, err := database.OpenDB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	// GORM veritabanı bağlantısını aç (gorm.DB)
	dsn := "bursaweb_ajans:Genetik1997.*/@tcp(84.54.13.3:3306)/bursaweb_crm?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	err = gormDB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	// UserController'ı oluştur
	userController := controller.NewUserControllerWithDB(sqlDB, gormDB)

	// Gin rota motorunu başlat
	r := gin.Default()

	// Rotaları ayarla
	routes.SetupRoutes(r, userController)

	// Uygulamayı belirli bir bağlantı noktasında başlat
	err = r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
