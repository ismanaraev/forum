package handler

import (
	"fmt"
	"forum3/internal/models"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

func (h *Handler) postPage(w http.ResponseWriter, r *http.Request) {
	// email := r.Context().Value("email")
}

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/create-post" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		html, err := template.ParseFiles("../internal/template/html/createPost.html")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		html.Execute(w, nil)

	case http.MethodPost:

		uuidString := r.Context().Value("uuid")
		value := uuidString.(string)

		uuid, err := uuid.FromString(value)
		if err != nil {
			return
		}

		data := &models.Auth{
			Uuid: uuid,
		}

		userInfo, err := h.service.GetUsersInfoByUUIDtoRepo(data.Uuid)
		if err != nil {
			log.Fatalf("Get user info from handler don`t work %e", err)
		}

		r.ParseForm()

		title, ok := r.Form["title"]
		if !ok {
			http.Error(w, "title field not found", http.StatusInternalServerError)
			return
		}

		contentR, ok := r.Form["content"]
		if !ok {
			http.Error(w, "content field not found", http.StatusInternalServerError)
			return
		}
		content := strings.Join(contentR, " ")

		category, ok := r.Form["chooseCategory"]
		if !ok {
			http.Error(w, "category field not found", http.StatusInternalServerError)
			return
		}
		categoryStr := strings.Join(category, ",")

		postData := &models.Post{
			Uuid:       uuid,
			Title:      title[0],
			Content:    content,
			Author:     userInfo.Username,
			CreatedAt:  time.Now().Format("yyyy-mm-dd HH:mm:ss"),
			Categories: categoryStr,
		}

		status, err := h.service.CreatePostService(*postData)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(status), status)
			log.Printf("Post not created")
		}

		http.Redirect(w, r, "/homepage", http.StatusSeeOther)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {}
