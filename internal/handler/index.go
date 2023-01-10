package handler

import (
	"forumv2/internal/models"
	"log"
	"net/http"
	"text/template"
)

type AllData struct {
	Data models.User
	Post []models.Post
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	html, err := template.ParseFiles("../internal/template/html/index.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	token, err := r.Cookie("session_name")
	if err != nil {
		category := r.FormValue("category")
		postData, err := h.service.GetAllPostService(category)
		if err != nil {
			log.Fatalf("Get all post from handler don`t work %e", err)
		}

		result := &AllData{
			Post: postData,
		}

		html.Execute(w, result)
		return

	} else if r.Method == http.MethodGet {
		html, err := template.ParseFiles("../internal/template/html/index.html")
		data := &models.User{}
		// по токену запрашиваем почту пользователя
		data.Uuid, err = h.service.GetSessionService(token.Value)
		if err != nil {
			http.Redirect(w, r, "/sign-up", http.StatusSeeOther)
			return
		}
		userInfo, err := h.service.GetUsersInfoByUUIDService(data.Uuid)
		if err != nil {
			log.Fatalf("Get user info from handler don`t work %e", err)
		}
		category := r.FormValue("category")
		postData, err := h.service.GetAllPostService(category)
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
