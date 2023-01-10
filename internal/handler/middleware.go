package handler

import (
	"context"
	"log"
	"net/http"
)

func (h *Handler) IsAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("session_name")
		if err != nil {
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}
		// по токену запрашиваем uuid пользователя
		uuid, err := h.service.GetSessionService(token.Value)
		if err != nil {
			log.Fatalf("Get session from handler don`t work %e", err)
		}
		ctx := context.WithValue(r.Context(), "uuid", uuid)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
