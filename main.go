package main

import (
	"github.com/gin-gonic/gin"
	"crm/database"
	"crm/routes"
)

func main() {
	db, err := database.OpenDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()

	routes.SetupRoutes(r) 

	r.Run(":3000")
}
