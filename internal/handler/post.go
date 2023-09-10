package handler

import (
	"database/sql"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"forumv2/internal/models"
)

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	postid := strings.TrimPrefix(r.URL.Path, PostAddress)
	postID, err := strconv.ParseInt(postid, 10, 64)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	post, err := h.service.GetPostByIDinService(models.PostID(postID))
	if err != nil {
		log.Print(err)
		if errors.Is(err, sql.ErrNoRows) {
			errorHeader(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	comments, err := h.service.GetCommentsByIDinService(post.ID)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles(TemplateDir + "html/commentPage.html")
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	idCtx := r.Context().Value("UserID")
	if idCtx == nil {
		res := struct {
			User    models.User
			Post    models.Post
			Comment []models.Comment
		}{User: models.User{}, Post: post, Comment: comments}
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
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	res := struct {
		User    models.User
		Post    models.Post
		Comment []models.Comment
	}{User: user, Post: post, Comment: comments}
	err = tmpl.Execute(w, &res)
	if err != nil {
		log.Print(err)
	}
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles(TemplateDir + "html/createPost.html")
		if err != nil {
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		idCtx := r.Context().Value("UserID")
		if idCtx == nil {
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		id := idCtx.(models.UserID)
		user, err := h.service.GetUsersInfoByUUIDService(id)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err = r.ParseForm()
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		title := r.PostFormValue("title")
		if title == "" {
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		content := r.PostFormValue("content")
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		categoriesArr := r.PostForm["categories"]
		for _, val := range categoriesArr {
			err = h.service.CreateCategory(val)
			if err != nil {
				log.Print(err)
				errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
		}
		post := models.Post{
			Title:     title,
			Content:   content,
			Author:    user,
			CreatedAt: time.Now(),
			Like:      0,
			Dislike:   0,
		}
		err = h.service.CheckPostInput(post)
		if err != nil {
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		postId, err := h.service.CreatePostService(post)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, PostAddress+strconv.FormatInt(int64(postId), 10), http.StatusSeeOther)
	default:
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
