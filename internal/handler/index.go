package handler

import (
	"forumv2/internal/models"
	"html/template"
	"log"
	"net/http"
)

type AllData struct {
	Data models.User
	Post []models.Post
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHeader(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	html, err := template.ParseFiles(TemplateDir + "html/index.html")
	if err != nil {
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res, err := h.service.GetAllPostService()
	if err != nil {
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	idCtx := r.Context().Value("UserID")
	if idCtx == nil {
		result := &AllData{
			Post: res,
		}

		err = html.Execute(w, result)
		if err != nil {
			log.Print(err)
		}
		return
	}
	id := idCtx.(models.UserID)

	userInfo, err := h.service.GetUsersInfoByUUIDService(id)
	if err != nil {
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	postData, err := h.service.GetAllPostService()
	if err != nil {
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result := &AllData{
		Data: userInfo,
		Post: postData,
	}
	err = html.Execute(w, result)
	if err != nil {
		log.Print(err)
	}
}
