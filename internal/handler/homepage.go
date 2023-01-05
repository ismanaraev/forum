package handler

import (
	"forum3/internal/models"
	"html/template"
	"log"
	"net/http"
)

type AllData struct {
	Data models.Auth
	Post []models.Post
}

func (h *Handler) homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/homepage" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	html, err := template.ParseFiles("../internal/template/html/homepage.html")
	if err != nil {

		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	token, err := r.Cookie("session_name")

	if err != nil {
		postData, err := h.service.GetAllPostService()
		if err != nil {
			log.Fatalf("Get all post from handler don`t work %e", err)
		}

		result := &AllData{
			Post: postData,
		}

		html.Execute(w, result)

	} else if r.Method == http.MethodGet {

		html, err := template.ParseFiles("../internal/template/html/homepage.html")
		data := &models.Auth{}

		// по токену запрашиваем почту пользователя
		data.Uuid, err = h.service.GetSessionRQtoRepo(token.Value)
		if err != nil {
			http.Redirect(w, r, "/sign-up", http.StatusSeeOther)
			return
		}

		userInfo, err := h.service.GetUsersInfoByUUIDtoRepo(data.Uuid)
		if err != nil {
			log.Fatalf("Get user info from handler don`t work %e", err)
		}

		postData, err := h.service.GetAllPostService()
		if err != nil {
			log.Fatalf("Get all post from handler don`t work %e", err)
		}

		result := &AllData{
			Data: userInfo,
			Post: postData,
		}

		html.Execute(w, result)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}
}
