package controllers

import (
	"blog/admin/helpers"
	"blog/admin/models"
	"fmt"
	"html/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type Contact struct{}

func (contact Contact) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.New("index").Funcs(template.FuncMap{}).ParseFiles(helpers.Include("contact/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Contacts"] = models.Contact{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (contact Contact) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	data := models.Contact{}.Get(params.ByName("id"))
	data.Delete()
	http.Redirect(w, r, "/admin/contact", http.StatusSeeOther)
}
