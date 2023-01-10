package handler

import (
	"fmt"
	"forumv2/internal/models"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type AllDataComments struct {
	Post         models.Post
	Comment      []models.Comment
	LikePost     models.LikePost
	LikeComments models.LikeComment
}

func (h *Handler) comment(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.URL.Path, "/comments/") {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	html, err := template.ParseFiles("../internal/template/html/commentPage.html")
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
		data := &models.User{}
		data.Uuid, err = h.service.GetSessionService(token.Value)
		if err != nil {
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}
		userInfo, err := h.service.GetUsersInfoByUUIDService(data.Uuid)
		if err != nil {
			log.Fatalf("Get user info from handler don`t work %e", err)
		}
		comments := r.FormValue("commentsS")
		if comments != "" {
			com := models.Comment{
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
		allReaction := models.LikePost{
			UserID: data.Uuid,
			PostID: postID,
			Status: vote,
		}
		likeReaction, err := h.service.LikePostService(allReaction)
		if err != nil {
			log.Printf("likeReaction error: %v", err)
			http.Error(w, "a", http.StatusInternalServerError)
		}
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
