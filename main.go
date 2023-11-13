package main

import (
	admin_models "blog/admin/models"
	"blog/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()
	admin_models.User{}.Migrate()
	admin_models.Category{}.Migrate()
	admin_models.Optik{}.Migrate()
	admin_models.Optikokuyucu{}.Migrate()
	admin_models.Dosya{}.Migrate()
	admin_models.General_Setting{}.Migrate()
	http.ListenAndServe(":8080", config.Routes())
}
