package handler

import (
	"forumv2/internal/models"
	"net/http"
	"text/template"

	"github.com/gofrs/uuid"
)

type AllData struct {
	Data models.User
	Post []models.Post
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHeader(w, http.StatusNotFound)
		//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		errorHeader(w, http.StatusMethodNotAllowed)
		//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	html, err := template.ParseFiles(TemplateDir + "html/index.html")
	if err != nil {
		errorHeader(w, http.StatusNotFound)

		//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	res, err := h.service.GetAllPostService()
	if err != nil {
		errorHeader(w, http.StatusNotFound)

		//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	uuidCtx := r.Context().Value("uuid")
	if uuidCtx == nil {
		result := &AllData{
			Post: res,
		}

		html.Execute(w, result)
		return
	}
	uuid := uuidCtx.(uuid.UUID)

	userInfo, err := h.service.GetUsersInfoByUUIDService(uuid)
	if err != nil {
		errorHeader(w, http.StatusInternalServerError)

		//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	postData, err := h.service.GetAllPostService()
	if err != nil {
		errorHeader(w, http.StatusInternalServerError)

		//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result := &AllData{
		Data: userInfo,
		Post: postData,
	}
	html.Execute(w, result)
}
