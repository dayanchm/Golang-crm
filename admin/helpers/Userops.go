package helpers

import (
	"blog/admin/models"
	"net/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetUser(w http.ResponseWriter, r *http.Request, username string, password string) error {
	session, err := store.Get(r, "blog-user")
	if err != nil {
		return err
	}
	session.Values["username"] = username
	session.Values["password"] = password

	return session.Save(r, w)
}

func CheckUser(w http.ResponseWriter, r *http.Request) bool {
	session, err := store.Get(r, "blog-user")
	if err != nil {
		return false
	}

	username := session.Values["username"]
	password := session.Values["password"]

	db, err := gorm.Open(mysql.Open(models.Dns), &gorm.Config{})
	if err != nil {
		SetAlert(w, r, "Database connection error")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return false
	}

	user, err := models.User{}.Get(db, "username = ? AND password = ?", username, password)
	if err != nil {
		SetAlert(w, r, "Please Login")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return false
	}

	if user.Username == username && user.Password == password {
		return true
	}

	SetAlert(w, r, "Please Login")
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	return false
}

func RemoveUser(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "blog-user")
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1

	return session.Save(r, w)
}
