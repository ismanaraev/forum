package handler

import (
	"forumv2/internal/models"
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) FilterByCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	categoriesArr, ok := r.URL.Query()["category"]
	if !ok {
		errorHeader(w, "invalid category", http.StatusBadRequest)
		return
	}
	posts, err := h.service.FilterPostsByCategories(categoriesArr)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	categories, err := h.service.GetAllCategories()
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles(TemplateDir + "html/index.html")
	if err != nil {
		log.Print(err)
		errorHeader(w, "", http.StatusInternalServerError)
		return
	}
	idCtx := r.Context().Value(MiddlewareUID)
	if idCtx == nil {
		res := AllData{
			Post:       posts,
			Categories: categories,
		}
		err = tmpl.Execute(w, &res)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		return
	}

	id := idCtx.(models.UserID)
	user, err := h.service.GetUsersInfoByUUIDService(id)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	res := AllData{
		Data:       user,
		Post:       posts,
		Categories: categories,
	}
	err = tmpl.Execute(w, &res)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		errorHeader(w, err.Error(), http.StatusInternalServerError)
		return
	}
	catArr, ok := r.PostForm["name"]
	if !ok {
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = h.service.CreateCategory(catArr[0])
	if err != nil {
		errorHeader(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
