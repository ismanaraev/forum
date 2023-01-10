package handler

import (
	"forumv2/internal/models"
	"net/http"
	"text/template"

	"github.com/gofrs/uuid"
)

type Data struct {
	Userinfo models.User
	Post     []models.Post
	LikePost []models.Post
}

func (h *Handler) myprofile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/myprofile" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	uuidString := r.Context().Value("uuid")
	value := uuidString.(string)

	uuid, err := uuid.FromString(value)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		html, err := template.ParseFiles(TemplateDir + "html/myprofile.html")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		userInfo, err := h.service.GetUsersInfoByUUIDService(uuid)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		usersPost, err := h.service.GetUsersPostInService(uuid)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		userLikePosts, err := h.service.GetUserLikePostsInService(uuid)

		data := &Data{
			Userinfo: userInfo,
			Post:     usersPost,
			LikePost: userLikePosts,
		}

		html.Execute(w, data)

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
