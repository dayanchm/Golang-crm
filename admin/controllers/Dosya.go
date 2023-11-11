package controllers

import (
	"blog/admin/helpers"
	"blog/admin/models"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
)

type Dosya struct{}

func (dosya Dosya) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.New("index").Funcs(template.FuncMap{}).ParseFiles(helpers.Include("dosya/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Dosya"] = models.Dosya{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (dosya Dosya) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("/dosya/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Dosya"] = models.Dosya{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (dosya Dosya) UploadDosya(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	title := r.FormValue("dosya-title")
	slug := slug.Make(title)
	description := r.FormValue("dosya-desc")
	content := r.FormValue("dosya-content")

	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println(err)
		return
	}
	models.Dosya{
		DosyaTitle:       title,
		DosyaSlug:        slug,
		DosyaDescription: description,
		DosyaContent:     content,
		Dosya_Url:        "uploads/" + header.Filename,
	}.Add()
	helpers.SetAlert(w, r, "Content-Added")
	http.Redirect(w, r, "/admin/dosya", http.StatusSeeOther)
}

func (dosya Dosya) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	data := models.Dosya{}.Get(params.ByName("id"))
	data.Delete()
	http.Redirect(w, r, "/admin/dosya", http.StatusSeeOther)
}

func (dosya Dosya) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("dosya/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Dosya"] = models.Dosya{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}
