package handler

import (
	"forumv2/internal/models"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gofrs/uuid"
)

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	postid := strings.TrimPrefix(r.URL.Path, PostAddress)
	postID, err := strconv.ParseInt(postid, 10, 64)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	post, err := h.service.GetPostByIDinService(postID)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	comments, err := h.service.GetCommentsByIDinService(post.ID)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles(TemplateDir + "html/commentPage.html")
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	res := struct {
		Post    models.Post
		Comment []models.Comment
	}{Post: post, Comment: comments}
	err = tmpl.Execute(w, &res)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles(TemplateDir + "html/createPost.html")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)

	case http.MethodPost:
		uuidCtx := r.Context().Value("uuid")
		if uuidCtx == nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		uuid := uuidCtx.(uuid.UUID)
		user, err := h.service.GetUsersInfoByUUIDService(uuid)
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err = r.ParseForm()
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		title := r.PostFormValue("title")
		if title == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		content := r.PostFormValue("content")
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		categoriesArr, ok := r.PostForm["categories"]
		if !ok {
			categoriesArr = []string{"other"}
		}
		categories, err := h.service.CreateCategory(categoriesArr)
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		post := models.Post{
			Uuid:       uuid,
			Title:      title,
			Content:    content,
			Author:     user.Username,
			CreatedAt:  time.Now().String(),
			Categories: categories,
			Like:       0,
			Dislike:    0,
		}
		id, err := h.service.CreatePostService(post)
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, PostAddress+strconv.FormatInt(id, 10), http.StatusSeeOther)
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return

	}
}
