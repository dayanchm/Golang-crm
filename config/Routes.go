package config

import (
	admin "blog/admin/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	// Dashboard
	r.GET("/admin", admin.Dashboard{}.Dashboard)

	// General
	r.GET("/admin/setting", admin.General{}.Index)
	r.POST("/admin/setting/", admin.General{}.Add)

	// Register
	r.GET("/admin/users_register", admin.Userops{}.Register)
	r.POST("/admin/register_list", admin.Userops{}.RegisterAdd)
	r.GET("/admin/register_list", admin.Userops{}.RegisterList)

	// Userops
	r.GET("/admin/login", admin.Userops{}.Index)
	r.POST("/admin/do_login", admin.Userops{}.Login)
	r.GET("/admin/logout", admin.Userops{}.Logout)

	// Serve File
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/assets/*filepath", http.Dir("site/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	r.ServeFiles("/admin/optikuploads/*filepath", http.Dir("optikuploads"))

	return r
}
