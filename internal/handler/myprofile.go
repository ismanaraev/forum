package handler

import (
	"forumv2/internal/models"
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Data     models.User
	Post     []models.Post
	LikePost []models.Post
}

func (h *Handler) myprofile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/myprofile" {
		errorHeader(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	html, err := template.ParseFiles(TemplateDir+"html/myprofile.html", TemplateDir+"html/header.html")
	if err != nil {
		errorHeader(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	idCtx := r.Context().Value(MiddlewareUID)
	id := idCtx.(models.UserID)

	switch r.Method {
	case http.MethodGet:
		userInfo, err := h.service.GetUsersInfoByUUIDService(id)
		if err != nil {
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		usersPost, err := h.service.GetUsersPostInService(id)
		if err != nil {
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		userLikePosts, err := h.service.GetUsersLikedPosts(id)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data := &Data{
			Data:     userInfo,
			Post:     usersPost,
			LikePost: userLikePosts,
		}

		err = html.Execute(w, data)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

	default:
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
