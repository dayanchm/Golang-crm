package controllers

import (
	"blog/admin/helpers"
	"blog/admin/models"
	"crypto/sha256"
	"fmt"
	"html/template"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Userops struct{}

func (userops Userops) sanitizeInput(input string) string {
	return url.QueryEscape(input)
}

func (userops Userops) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("userops/login")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	alert := helpers.GetAlert(w, r)
	data["is_alert"] = alert["is_alert"]
	data["message"] = alert["message"]
	csrfToken, err := helpers.SetCSRFToken(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	data["CSRFToken"] = csrfToken
	view.ExecuteTemplate(w, "index", data)
}
func (userops Userops) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	allowed, remainingTime := helpers.CheckLoginAttempts(r)
	if !allowed {
		// Dakika ve saniye cinsinden kalan süreyi hesaplayın
		minutes := int(remainingTime.Minutes())
		seconds := int(remainingTime.Seconds()) % 60
		helpers.SetAlert(w, r, fmt.Sprintf("Too many login attempts. Please try again after %d minutes and %d seconds.", minutes, seconds))
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	username := userops.sanitizeInput(r.FormValue("username"))
	password := userops.sanitizeInput(r.FormValue("password"))

	if username == "" || password == "" {
		helpers.SetAlert(w, r, "Username and Password cannot be empty")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	hashedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	db, err := gorm.Open(mysql.Open(models.Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		helpers.SetAlert(w, r, "Database connection error")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	user, err := models.User{}.Get(db, "username = ? AND password = ?", username, hashedPassword)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.SetAlert(w, r, "Wrong Username and Password")
		} else {
			helpers.SetAlert(w, r, "User not found")
		}
		helpers.RecordLoginAttempt(r, false)
		fmt.Println(err)
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	if user.Username == username && user.Password == hashedPassword {
		helpers.SetUser(w, r, username, hashedPassword)
		helpers.RecordLoginAttempt(r, true)
		helpers.SetAlert(w, r, "Welcome")
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		helpers.RecordLoginAttempt(r, false)
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
	db, err := gorm.Open(mysql.Open(models.Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	user, err := models.User{}.Get(db)
	data["User"] = user
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
	db, err := gorm.Open(mysql.Open(models.Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	user, err := models.User{}.Get(db)
	data["User"] = user
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
	if err := userModel.AddRole(db, &newRole); err != nil {
		fmt.Println(err)
		return
	}

	helpers.SetAlert(w, r, "User registered successfully")
	http.Redirect(w, r, "/admin/register_list", http.StatusSeeOther)
}
func (userops Userops) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	helpers.RemoveUser(w, r)
	helpers.SetAlert(w, r, "Good bye, see you soon")
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}
