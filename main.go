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
	gin.SetMode(gin.ReleaseMode)

	sqlDB, err := database.OpenDB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	dsn := "bursaweb_ajans:Genetik1997.*/@tcp(84.54.13.3:3306)/bursaweb_crm?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	err = gormDB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	userController := controller.NewUserControllerWithDB(sqlDB, gormDB)

	router := gin.Default()

	routes.SetupRoutes(router, userController)

	err = router.Run(":3002")
	if err != nil {
		panic(err)
	}
}
