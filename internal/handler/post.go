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
		errorHeader(w, "", http.StatusMethodNotAllowed)
		//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	postid := strings.TrimPrefix(r.URL.Path, PostAddress)
	postID, err := strconv.ParseInt(postid, 10, 64)
	if err != nil {
		log.Print(err)
		errorHeader(w, "", http.StatusInternalServerError)
		//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	post, err := h.service.GetPostByIDinService(postID)
	if err != nil {
		log.Print(err)
		errorHeader(w, "", http.StatusInternalServerError)
		//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	comments, err := h.service.GetCommentsByIDinService(post.ID)
	if err != nil {
		log.Print(err)
		errorHeader(w, "", http.StatusInternalServerError)
		//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles(TemplateDir + "html/commentPage.html")
	if err != nil {
		log.Print(err)
		errorHeader(w, "", http.StatusInternalServerError)
		//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	res := struct {
		Post    models.Post
		Comment []models.Comment
	}{Post: post, Comment: comments}
	err = tmpl.Execute(w, &res)
	if err != nil {
		log.Print(err)
		errorHeader(w, "", http.StatusInternalServerError)
		//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles(TemplateDir + "html/createPost.html")
		if err != nil {
			errorHeader(w, "", http.StatusInternalServerError)
			//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)

	case http.MethodPost:
		uuidCtx := r.Context().Value("uuid")
		if uuidCtx == nil {
			errorHeader(w, "", http.StatusBadRequest)
			//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		uuid := uuidCtx.(uuid.UUID)
		user, err := h.service.GetUsersInfoByUUIDService(uuid)
		if err != nil {
			log.Print(err)
			errorHeader(w, "", http.StatusInternalServerError)
			//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err = r.ParseForm()
		if err != nil {
			log.Print(err)
			errorHeader(w, "", http.StatusBadRequest)
			//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		title := r.PostFormValue("title")
		if title == "" {
			errorHeader(w, "", http.StatusBadRequest)
			//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		content := r.PostFormValue("content")
		if err != nil {
			log.Print(err)
			errorHeader(w, "", http.StatusBadRequest)
			//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		categoriesArr := r.PostForm["categories"]
		categories, err := h.service.CreateCategory(categoriesArr)
		if err != nil {
			log.Print(err)
			errorHeader(w, "", http.StatusBadRequest)
			//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		t := time.Now()
		timeFormat := t.Format("15:04:04,02 January 2006")
		post := models.Post{
			Uuid:       uuid,
			Title:      title,
			Content:    content,
			Author:     user.Username,
			CreatedAt:  timeFormat,
			Categories: categories,
			Like:       0,
			Dislike:    0,
		}
		err = h.service.CheckPostInput(post)
		if err != nil {
			errorHeader(w, "Post Is Filled Out Incorrectly", http.StatusBadRequest)
			//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		id, err := h.service.CreatePostService(post)
		if err != nil {
			log.Print(err)
			errorHeader(w, "", http.StatusInternalServerError)
			//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, PostAddress+strconv.FormatInt(id, 10), http.StatusSeeOther)
	default:
		errorHeader(w, "", http.StatusBadRequest)
		//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return

	}
}
