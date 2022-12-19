package handler

import (
	"forum3/internal/models"
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/homepage" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodPost:

	case http.MethodGet:
		data := &models.Auth{}
		token, err := r.Cookie("session_name")
		if err != nil {
			html, err := template.ParseFiles("../internal/template/html/homepage.html")
			if err != nil {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

			// data, err := h.service.GetAllPostService()
			// if err != nil {
			// 	log.Fatalf("Get all post from handler don`t work %e", err)
			// }

			html.Execute(w, nil)

		} else {
			html, err := template.ParseFiles("../internal/template/html/homepage.html")
			if err != nil {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

			// по токену запрашиваем почту пользователя
			data.Uuid, err = h.service.GetSessionRQtoRepo(token.Value)
			if err != nil {
				log.Fatalf("Get session from handler don`t work %e", err)
			}

			userInfo, err := h.service.GetUsersInfoByUUIDtoRepo(data.Uuid)
			if err != nil {
				log.Fatalf("Get user info from handler don`t work %e", err)
			}

			html.Execute(w, userInfo)
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
