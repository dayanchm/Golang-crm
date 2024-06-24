package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
)

func GenerateCSRFToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func SetCSRFToken(w http.ResponseWriter, r *http.Request) (string, error) {
	token, err := GenerateCSRFToken()
	if err != nil {
		return "", err
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "csrf_token",
		Value: token,
		Path:  "/",
	})
	return token, nil
}

func ValidateCSRFToken(r *http.Request) bool {
	cookie, err := r.Cookie("csrf_token")
	if err != nil {
		return false
	}
	token := r.FormValue("csrf_token")
	return cookie.Value == token
}
