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

type Slider struct {
}

func (slider Slider) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.New("index").Funcs(template.FuncMap{}).ParseFiles(helpers.Include("/slider/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Sliders"] = models.Slider{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (slider Slider) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("/slider/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Sliders"] = models.Slider{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (slider Slider) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	title := r.FormValue("slider-title")
	slug := slug.Make(title)

	// upload-picture
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("slider-picture")
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
	models.Slider{
		Slider_Title:       title,
		Slider_Slug:        slug,
		Slider_Picture_url: "optikuploads/" + header.Filename,
	}.Add()
	helpers.SetAlert(w, r, "Eklendi")
	http.Redirect(w, r, "/admin/slider", http.StatusSeeOther)
}

func (slider Slider) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	data := models.Slider{}.Get(params.ByName("id"))
	data.Delete()
	http.Redirect(w, r, "/admin/slider", http.StatusSeeOther)
}

func (slider Slider) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("slider/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Slider"] = models.Slider{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}

func (slider Slider) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	data := models.Slider{}.Get(params.ByName("id"))
	title := r.FormValue("slider-title")
	slug := slug.Make(title)
	slideris_selected := r.FormValue("slider_selected")
	var picture_url string

	if slideris_selected == "1" {
		// upload-picture
		r.ParseMultipartForm(10 << 20)
		file, header, err := r.FormFile("slider-picture")
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
		os.Remove(data.Slider_Picture_url)
	} else {
		picture_url = data.Slider_Picture_url
	}

	data.Updates(models.Slider{
		Slider_Title:       title,
		Slider_Slug:        slug,
		Slider_Picture_url: picture_url,
	})
	http.Redirect(w, r, "/admin/slider/edit/"+params.ByName("id"), http.StatusSeeOther)
}
