package handler

import (
	"context"
	"net/http"
)

func (h *Handler) IsAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("session_token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			http.Redirect(w, r, "/sign-in", http.StatusNotFound)
			// io.WriteString(w, "invalid token")
			// return
		}

		// по токену запрашиваем почту пользователя
		uuid, err := h.service.GetSessionRQtoRepo(token.Value)
		ctx := context.WithValue(r.Context(), "uuid", uuid)
		next.ServeHTTP(w, r.WithContext(ctx))

		// Продлеваем жизнь токена
		cookie := http.Cookie{
			SameSite: http.SameSiteNoneMode,
			Name:     "session_token",
			Value:    token.Value,
			MaxAge:   3600 * 12,
		}
		http.SetCookie(w, &cookie)
	}
}
