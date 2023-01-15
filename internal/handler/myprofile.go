package handler

import (
	"forumv2/internal/models"
	"html/template"
	"net/http"

	"github.com/gofrs/uuid"
)

type Data struct {
	Userinfo models.User
	Post     []models.Post
	LikePost []models.Post
}

func (h *Handler) myprofile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/myprofile" {
		errorHeader(w, "", http.StatusNotFound)
		//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	html, err := template.ParseFiles(TemplateDir + "html/myprofile.html")
	if err != nil {
		errorHeader(w, "", http.StatusNotFound)

		//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	uuidString := r.Context().Value("uuid")
	uuid := uuidString.(uuid.UUID)

	switch r.Method {
	case http.MethodGet:

		userInfo, err := h.service.GetUsersInfoByUUIDService(uuid)
		if err != nil {
			errorHeader(w, "", http.StatusInternalServerError)
			//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		usersPost, err := h.service.GetUsersPostInService(uuid)
		if err != nil {
			errorHeader(w, "", http.StatusInternalServerError)
			//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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
		errorHeader(w, "", http.StatusMethodNotAllowed)

		//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
