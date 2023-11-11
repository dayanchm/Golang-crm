package config

import (
	admin "blog/admin/controllers"
	site "blog/site/controllers"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	// Dashboard
	r.GET("/dashboard", admin.Dashboard{}.Dashboard)
	//Add-BLOG
	r.GET("/admin/blog", admin.Dashboard{}.Index)
	r.GET("/admin/content-add", admin.Dashboard{}.NewItem)
	r.POST("/admin/blog", admin.Dashboard{}.Add)
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	r.POST("/admin/edit/:id", admin.Dashboard{}.Update)

	// Admin-Optikform
	r.GET("/admin/optik", admin.Optik{}.Index)
	r.GET("/optik/add", admin.Optik{}.NewItem)
	r.POST("/admin/optik", admin.Optik{}.Add)
	r.GET("/admin/optik/edit/:id", admin.Optik{}.Edit)
	r.POST("/optik/update/:id", admin.Optik{}.Update)
	r.GET("/optik/delete/:id", admin.Optik{}.Delete)

	// Admin-OptikOkuyucu
	r.GET("/admin/optikokuyucu", admin.Optikokuyucu{}.Index)
	r.GET("/optikokuyucu/add", admin.Optikokuyucu{}.NewItem)
	r.POST("/admin/optikokuyucu", admin.Optikokuyucu{}.Add)
	r.GET("/admin/optikokuyucu/edit/:id", admin.Optikokuyucu{}.Edit)
	r.POST("/optikokuyucu/update/:id", admin.Optikokuyucu{}.Update)
	r.GET("/optikokuyucu/delete/:id", admin.Optikokuyucu{}.Delete)

	//Admin-Slider
	r.GET("/admin/slider", admin.Slider{}.Index)
	r.GET("/admin/slider-add", admin.Slider{}.NewItem)
	r.POST("/admin/slidercontent", admin.Slider{}.Add)
	r.GET("/admin/slider/edit/:id", admin.Slider{}.Edit)
	r.GET("/slider/delete/:id", admin.Slider{}.Delete)
	r.POST("/slider/update/:id", admin.Slider{}.Update)

	//Admin-Dosya Kaydet
	r.GET("/admin/registers", admin.User{}.Index)
	r.POST("/admin/registersadd", admin.User{}.Add)
	r.GET("/admin/registersadd", admin.User{}.NewItem)
	r.PUT("/admin/registers/edit/:id", admin.User{}.Edit)
	r.DELETE("/registers/delete/:id", admin.User{}.Delete)

	// Kayit ol
	r.GET("/admin/dosya", admin.Userops{}.Index)
	r.GET("/admin/dosyaupload", admin.Dosya{}.NewItem)
	r.POST("/admin/dosyacontentupload", admin.Dosya{}.UploadDosya)
	r.GET("/admin/dosya/edit/:id", admin.Dosya{}.Edit)
	r.GET("/dosya/delete/:id", admin.Dosya{}.Delete)

	//Categories
	r.GET("/admin/kategoriler", admin.Categories{}.Index)
	r.POST("/admin/kategoriler/add", admin.Categories{}.Add)
	r.GET("/admin/kategoriler/delete/:id", admin.Categories{}.Delete)

	// İletisim
	r.GET("/admin/contact", admin.Contact{}.Index)
	r.GET("/contact/delete/:id", admin.Contact{}.Delete)

	//Userops
	r.GET("/admin/login", admin.Userops{}.Index)
	r.POST("/admin/do_login", admin.Userops{}.Login)
	r.GET("/admin/logout", admin.Userops{}.Logout)

	//SITE
	//Homepage
	r.GET("/", site.Homepage{}.Index)
	r.GET("/about", site.Homepage{}.About)
	r.GET("/optik", site.Homepage{}.Optik)
	r.GET("/optikformdetail/:optik_slug", site.Homepage{}.Optikdetail)
	r.GET("/optikokuyucu", site.Homepage{}.Optikokuyucu)
	r.GET("/optikokuyucudetail/:optik_slug", site.Homepage{}.Optikokuyucudetail)
	r.GET("/contact", site.Homepage{}.İletisim)
	r.POST("/iletisim/contact", site.Homepage{}.Post)
	r.GET("/blog", site.Homepage{}.Blog)
	r.GET("/blog/:slug", site.Homepage{}.Detail)
	r.GET("/dosyalar", site.Homepage{}.Dosyalar)
	r.GET("/dosyalar/:dosya_slug", site.Homepage{}.Dosyalarslug)
	r.GET("/thanks", site.Homepage{}.Thanks)

	//Serve File
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/assets/*filepath", http.Dir("site/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	r.ServeFiles("/admin/optikuploads/*filepath", http.Dir("optikuploads"))
	return r
}
