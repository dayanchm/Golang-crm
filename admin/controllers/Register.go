package controllers

import (
	"blog/admin/helpers"
	"blog/admin/models"
	"fmt"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type User struct {
}

func (user User) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.New("index").Funcs(template.FuncMap{}).ParseFiles(helpers.Include("/register/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["User"] = models.User{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (user User) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("/register/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["User"] = models.User{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (user User) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	title := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	content := r.FormValue("confirm_password")
	models.User{
		Username:        title,
		Email:           email,
		Password:        password,
		ConfirmPassword: content,
	}.Add()
	helpers.SetAlert(w, r, "Eklendi")
	http.Redirect(w, r, "/admin/registers", http.StatusSeeOther)
}
func (user User) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("register/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["User"] = models.User{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}

func (user User) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	User := models.User{}.Get(params.ByName("id"))
	title := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	content := r.FormValue("confirm_password")

	User.Updates(models.User{
		Username:        title,
		Email:           email,
		Password:        password,
		ConfirmPassword: content,
	})
	http.Redirect(w, r, "/admin/register/edit/"+params.ByName("id"), http.StatusSeeOther)
}

func (user User) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	data := models.Optikokuyucu{}.Get(params.ByName("id"))
	data.Delete()
	http.Redirect(w, r, "/admin/optikokuyucu", http.StatusSeeOther)
}
