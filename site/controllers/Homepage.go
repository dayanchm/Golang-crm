package controllers

import (
	"blog/site/helpers"
	"blog/site/models"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Homepage struct{}

func (homepage Homepage) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory": func(categoryID int) string {
			return models.Category{}.Get(categoryID).Title
		},
		"getDate": func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d", t.Day(), int(t.Month()), t.Year())
		},
	}).ParseFiles(helpers.Include("homepage/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Sliders"] = models.Slider{}.GetAll()
	view.ExecuteTemplate(w, "index", data)

}
func (homepage Homepage) Detail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("homepage/detail")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Post"] = models.Post{}.Get("slug = ?", params.ByName("slug"))
	view.ExecuteTemplate(w, "index", data)
}
func (homepage Homepage) Optik(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.New("index").ParseFiles(helpers.Include("homepage/optik")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Optiks"] = models.Optik{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}
func (homepage Homepage) Optikdetail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("homepage/optikdetail")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Optik"] = models.Optik{}.Get("optik_slug = ?", params.ByName("optik_slug"))
	view.ExecuteTemplate(w, "index", data)
}
func (homepage Homepage) Optikokuyucu(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.New("index").ParseFiles(helpers.Include("homepage/optikokuyucu")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Optiksokuyucu"] = models.Optikokuyucu{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}
func (homepage Homepage) Optikokuyucudetail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("homepage/optikokuyucudetail")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Optiksokuyucu"] = models.Optikokuyucu{}.Get("optik_slug = ?", params.ByName("optik_slug"))
	view.ExecuteTemplate(w, "index", data)
}
func (homepage Homepage) İletisim(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.New("index").Funcs(template.FuncMap{}).ParseFiles(helpers.Include("homepage/iletisim")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}
func (homepage Homepage) Blog(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory": func(categoryID int) string {
			return models.Category{}.Get(categoryID).Title
		},
		"getDate": func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d", t.Day(), int(t.Month()), t.Year())
		},
	}).ParseFiles(helpers.Include("homepage/blog")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}
func (homepage Homepage) Dosyalar(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.New("index").Funcs(template.FuncMap{}).ParseFiles(helpers.Include("homepage/dosyalar")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Dosya"] = models.Dosya{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}
func (homepage Homepage) Dosyalarslug(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("homepage/dosyadetail")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Dosya"] = models.Dosya{}.Get("dosya_slug = ?", params.ByName("dosya_slug"))
	view.ExecuteTemplate(w, "index", data)
}
func (homepage Homepage) Post(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	adi := r.FormValue("adi")
	soyadi := r.FormValue("soyadi")
	email := r.FormValue("email")
	telefon := r.FormValue("phone")
	mesaj := r.FormValue("message")

	models.Contact{
		Adi:     adi,
		Soyadi:  soyadi,
		Email:   email,
		Telefon: telefon,
		Mesaj:   mesaj,
	}.Add()
	http.Redirect(w, r, "/thanks", http.StatusSeeOther)
}
func (homepage Homepage) Thanks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.New("index").Funcs(template.FuncMap{}).ParseFiles(helpers.Include("homepage/thanks")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	view.ExecuteTemplate(w, "index", data)

}
func (homepage Homepage) About(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.New("index").ParseFiles(helpers.Include("homepage/about")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Optiks"] = models.Optik{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}

