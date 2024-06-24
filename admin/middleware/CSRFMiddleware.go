package middleware

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"blog/admin/helpers"
)

func CSRFMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.Method == http.MethodPost {
			if !helpers.ValidateCSRFToken(r) {
				http.Error(w, "Invalid CSRF token", http.StatusForbidden)
				return
			}
		}
		next(w, r, ps)
	}
}

func Adapt(h http.Handler, fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h.ServeHTTP(w, r)
		fn(w, r, ps)
	}
}
