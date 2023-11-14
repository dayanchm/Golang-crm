package helpers

import (
	"fmt"
	"net/http"
)

func Opera(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")

	ipAddress := r.Header.Get("X-Forwarded-For")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Real-IP")
	}

	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
    <p>User-Agent: %s</p>
    <p>IP Adresi: %s</p>`, userAgent, ipAddress)

	if contains(userAgent, "Chrome") {
		fmt.Fprintln(w, "<p>Bu kullanıcı Chrome tarayıcısını kullanıyor.</p>")
	} else if contains(userAgent, "Firefox") {
		fmt.Fprintln(w, "<p>Bu kullanıcı Firefox tarayıcısını kullanıyor.</p>")
	} else {
		fmt.Fprintln(w, "<p>Bu kullanıcının tarayıcısı belirlenemedi.</p>")
	}
	fmt.Fprintln(w, ``)
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && (len(s)-len(substr)) >= 0 && s[len(s)-len(substr):] == substr
}
