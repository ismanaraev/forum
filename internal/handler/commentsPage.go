package handler

import (
	"fmt"
	"forum3/internal/models"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type AllDataComments struct {
	Post    models.Post
	Comment []models.Comments
}

func (h *Handler) comment(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.URL.Path, "/comments/") {

		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	html, err := template.ParseFiles("../internal/template/html/commentsPage.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	rx := regexp.MustCompile(`\w+$`)
	rawraq := r.URL.Path
	fr := rx.FindString(rawraq)
	postID, _ := strconv.Atoi(fr)

	switch r.Method {
	case http.MethodGet:

		content, err := h.service.GetPostByIDinService(postID)
		if err != nil {

			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		allComments, err := h.service.GetCommentsByIDinService(postID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		data := &AllDataComments{
			Post:    content,
			Comment: allComments,
		}

		html.Execute(w, data)
	case http.MethodPost:
		token, err := r.Cookie("session_name")
		if err != nil {
			http.Redirect(w, r, "/sign-up", http.StatusSeeOther)
			return
		}
		data := &models.Auth{}

		data.Uuid, err = h.service.GetSessionRQtoRepo(token.Value)
		if err != nil {
			http.Redirect(w, r, "/sign-up", http.StatusSeeOther)
			return
		}

		userInfo, err := h.service.GetUsersInfoByUUIDtoRepo(data.Uuid)
		if err != nil {
			log.Fatalf("Get user info from handler don`t work %e", err)
		}

		r.ParseForm()

		commentsR, ok := r.Form["commentsS"]
		if !ok {
			http.Error(w, "commentary field not found", http.StatusInternalServerError)
			return
		}
		comments := strings.Join(commentsR, " ")

		com := models.Comments{
			Author:    userInfo.Username,
			Content:   comments,
			CreatedAt: time.Now().Format(time.RFC1123),
			PostID:    postID,
		}

		status, err := h.service.CreateCommentsInService(com)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(status), status)
			log.Printf("Comment not created")
		}

		content, err := h.service.GetPostByIDinService(postID)
		if err != nil {

			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		allComments, err := h.service.GetCommentsByIDinService(postID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		result := &AllDataComments{
			Post:    content,
			Comment: allComments,
		}

		html.Execute(w, result)

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
