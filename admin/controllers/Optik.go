package controllers

import (
	"blog/admin/helpers"
	"blog/admin/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
)

type Optik struct {
}

func (optik Optik) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory": func(categoryID int) string {
			return models.Category{}.Get(categoryID).Title
		},
	}).ParseFiles(helpers.Include("/optik/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Optiks"] = models.Optik{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (optik Optik) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("/optik/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Optiks"] = models.Optik{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (optik Optik) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	title := r.FormValue("optik-title")
	slug := slug.Make(title)
	description := r.FormValue("optik-desc")
	content := r.FormValue("optik-content")

	// upload-picture
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("optik-picture")
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile("optikuploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println(err)
		return
	}
	models.Optik{
		OptikTitle:       title,
		OptikSlug:        slug,
		OptikDescription: description,
		OptikContent:     content,
		OptikPicture_url: "optikuploads/" + header.Filename,
	}.Add()
	helpers.SetAlert(w, r, "Eklendi")
	http.Redirect(w, r, "/admin/optik", http.StatusSeeOther)
}

func (optik Optik) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	data := models.Optik{}.Get(params.ByName("id"))
	data.Delete()
	http.Redirect(w, r, "/admin/optik", http.StatusSeeOther)
}

func (optik Optik) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("optik/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Optik"] = models.Optik{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}

func (optik Optik) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	Optik := models.Optik{}.Get(params.ByName("id"))
	title := r.FormValue("optik-title")
	slug := slug.Make(title)
	description := r.FormValue("optik-desc")
	content := r.FormValue("optik-content")
	is_selected := r.FormValue("optik_selected")
	var picture_url string

	if is_selected == "1" {
		// upload-picture
		r.ParseMultipartForm(10 << 20)
		file, header, err := r.FormFile("optik-picture")
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

		picture_url = "uploads/" + header.Filename
		os.Remove(Optik.OptikPicture_url)
	} else {
		picture_url = Optik.OptikPicture_url
	}

	Optik.Updates(models.Optik{
		OptikTitle:       title,
		OptikSlug:        slug,
		OptikDescription: description,
		OptikContent:     content,
		OptikPicture_url: picture_url,
	})
	http.Redirect(w, r, "/admin/optik/edit/"+params.ByName("id"), http.StatusSeeOther)
}
