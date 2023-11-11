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

type Optikokuyucu struct {
}

func (optik Optikokuyucu) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.New("index").Funcs(template.FuncMap{}).ParseFiles(helpers.Include("/optikokuyucu/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Optiksokuyucu"] = models.Optikokuyucu{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (optik Optikokuyucu) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("/optikokuyucu/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Optiksokuyucu"] = models.Optikokuyucu{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (optik Optikokuyucu) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	title := r.FormValue("optikokuyucu-title")
	slug := slug.Make(title)
	description := r.FormValue("optikokuyucu-desc")
	content := r.FormValue("optikokuyucu-content")

	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("optikokuyucu-picture")
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
	models.Optikokuyucu{
		OptikTitle:       title,
		OptikSlug:        slug,
		OptikDescription: description,
		OptikContent:     content,
		OptikPicture_url: "optikuploads/" + header.Filename,
	}.Add()
	helpers.SetAlert(w, r, "Eklendi")
	http.Redirect(w, r, "/admin/optikokuyucu", http.StatusSeeOther)
}
func (optik Optikokuyucu) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("optikokuyucu/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Optiksokuyucu"] = models.Optikokuyucu{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}

func (optik Optikokuyucu) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	Optikokuyucu := models.Optikokuyucu{}.Get(params.ByName("id"))
	title := r.FormValue("optikokuyucu-title")
	slug := slug.Make(title)
	description := r.FormValue("optikokuyucu-desc")
	content := r.FormValue("optikokuyucu-content")
	is_selected := r.FormValue("optikokuyucu_selected")
	var picture_url string

	if is_selected == "1" {
		// upload-picture
		r.ParseMultipartForm(10 << 20)
		file, header, err := r.FormFile("optikokuyucu-picture")
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
		os.Remove(Optikokuyucu.OptikPicture_url)
	} else {
		picture_url = Optikokuyucu.OptikPicture_url
	}

	Optikokuyucu.Updates(models.Optikokuyucu{
		OptikTitle:       title,
		OptikSlug:        slug,
		OptikDescription: description,
		OptikContent:     content,
		OptikPicture_url: picture_url,
	})
	http.Redirect(w, r, "/admin/optikokuyucu/edit/"+params.ByName("id"), http.StatusSeeOther)
}

func (optik Optikokuyucu) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	data := models.Optikokuyucu{}.Get(params.ByName("id"))
	data.Delete()
	http.Redirect(w, r, "/admin/optikokuyucu", http.StatusSeeOther)
}
