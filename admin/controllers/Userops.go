package controllers

import (
	"blog/admin/helpers"
	"blog/admin/models"
	"crypto/sha256"
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Userops struct{}

func (userops Userops) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("userops/login")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (userops Userops) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := r.FormValue("username")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password"))))
	user := models.User{}.Get("username = ? AND password = ?", username, password)
	if user.Username == username && user.Password == password {
		helpers.SetUser(w, r, username, password)
		helpers.SetAlert(w, r, "Welcome")
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		helpers.SetAlert(w, r, "Wrong Username and Password")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
}

func (userops Userops) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.New("index").Funcs(template.FuncMap{}).ParseFiles(helpers.Include("/register/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["User"] = models.User{}.Get()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (userops Userops) RegisterList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.New("index").Funcs(template.FuncMap{}).ParseFiles(helpers.Include("/register/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["User"] = models.User{}.Get()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (userops Userops) RegisterAdd(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	// Sayfa üzerindeki formdan rol adını al
	roleName := r.FormValue("role")

	// Diğer form verilerini al
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	contact := r.FormValue("contact")

	hashedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	// Yeni kullanıcı nesnesini oluştur
	userModel := models.User{
		Name:     name,
		Surname:  surname,
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Contact:  contact,
	}

	// Kullanıcının seçtiği rol adını kullanarak bir rol oluştur
	var newRole models.Role
	newRole.Name = roleName

	// Gorm DB bağlantısını al
	db, err := gorm.Open(mysql.Open(models.Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Kullanıcıya rolü ata
	userModel.AddRole(db, &newRole)

	helpers.SetAlert(w, r, "User registered successfully")
	http.Redirect(w, r, "/admin/register_list", http.StatusSeeOther)
}

func (userops Userops) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	helpers.RemoveUser(w, r)
	helpers.SetAlert(w, r, "Good bye, see you soon")
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}
