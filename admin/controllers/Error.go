package controllers

import (
	"net/http"
)

// Error is a struct that defines the Error controller.

type Error struct {
}

// NotFound is a method that handles the 404 error page.
func (e Error) NotFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin/views/error/index.html")
}
