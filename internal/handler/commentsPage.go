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
	Post         models.Post
	Comment      []models.Comments
	LikePost     models.LikePost
	LikeComments models.LikeComments
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

		result := &AllDataComments{
			Post:    content,
			Comment: allComments,
		}

		html.Execute(w, result)
	case http.MethodPost:
		token, err := r.Cookie("session_name")
		if err != nil {
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}
		data := &models.Auth{}

		data.Uuid, err = h.service.GetSessionRQtoRepo(token.Value)
		if err != nil {
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}

		userInfo, err := h.service.GetUsersInfoByUUIDtoRepo(data.Uuid)
		if err != nil {
			log.Fatalf("Get user info from handler don`t work %e", err)
		}

		// r.ParseForm()

		// commentsR, ok := r.Form["commentsS"]
		// if !ok {
		// 	http.Error(w, "commentary field not found", http.StatusInternalServerError)
		// 	return
		// }
		// comments := strings.Join(commentsR, " ")

		comments := r.FormValue("commentsS")
		if comments != "" {
			com := models.Comments{
				Author:    userInfo.Username,
				Content:   comments,
				CreatedAt: time.Now().Format(time.RFC1123),
				PostID:    postID,
			}

			status, err := h.service.CreateCommentsInService(com)
			if err != nil {
				http.Error(w, http.StatusText(status), status)
				log.Printf("Comment not created")
			}
		}

		content, err := h.service.GetPostByIDinService(postID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		allComments, err := h.service.GetCommentsByIDinService(postID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		var vote models.LikeStatus
		like := r.FormValue("like")
		dislike := r.FormValue("dislike")
		fmt.Println(like)
		fmt.Println(dislike)

		if like == "" {
			vote = models.DisLike
			dislike = ""
		} else {
			vote = models.Like
			like = ""
		}

		fmt.Println(vote)
		allReaction := models.LikePost{
			UserID: data.Uuid,
			PostID: postID,
			Status: vote,
		}

		likeReaction, err := h.service.LikeInService(allReaction)
		if err != nil {
			log.Printf("likeReaction error: %v", err)
			http.Error(w, "a", http.StatusInternalServerError)
		}

		// likeCounter := h.service.CounterLikeInService()

		result := &AllDataComments{
			Post:     content,
			Comment:  allComments,
			LikePost: likeReaction,
		}

		html.Execute(w, result)

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
