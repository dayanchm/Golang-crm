package controllers

import (
	"blog/admin/helpers"
	"blog/admin/models"
	"crypto/sha256"
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
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

/* func (userops Userops) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")
	email := r.FormValue("email")
	existingUser := models.User{}.Get("username = ? OR email = ?", username, email)
	if existingUser.Username == username {
		helpers.SetAlert(w, r, "Bu kullanıcı adı zaten alınmış")
		http.Redirect(w, r, "/admin/register", http.StatusSeeOther)
		return
	} else if existingUser.Email == email {
		helpers.SetAlert(w, r, "Bu e-posta zaten kullanılıyor")
		http.Redirect(w, r, "/admin/register", http.StatusSeeOther)
		return
	}
	if password != confirmPassword {
		helpers.SetAlert(w, r, "Şifreler eşleşmiyor")
		http.Redirect(w, r, "/admin/register", http.StatusSeeOther)
		return
	}
	hashedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	newUser := models.User{
		Username: username,
		Password: hashedPassword,
		Email:    email,
	}
	err := newUser.Create()
	if err != nil {
		helpers.SetAlert(w, r, "Kayıt sırasında bir hata oluştu. Lütfen daha sonra tekrar deneyin.")
		http.Redirect(w, r, "/admin/register", http.StatusSeeOther)
		return
	}
	helpers.SetAlert(w, r, "Kaydınız başarıyla tamamlandı. Lütfen giriş yapın.")
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
} */

func (userops Userops) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	helpers.RemoveUser(w, r)
	helpers.SetAlert(w, r, "Good bye, see you soon")
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}
