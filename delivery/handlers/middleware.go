package handlers

import (
	"net/http"
)

func (h Handler) IsAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h.service.User.CheckToken("ID", "TOKEN")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
