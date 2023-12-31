package handler

import (
	"forumv2/internal/models"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	idCtx := r.Context().Value(MiddlewareUID)
	if idCtx == nil {
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	id := idCtx.(models.UserID)
	user, err := h.service.GetUsersInfoByUUIDService(id)
	if err != nil {
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	postIDStr := r.FormValue("postID")
	if postIDStr == "" {
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	content := r.FormValue("content")
	if content == "" {
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	comm := models.Comment{
		PostID:    models.PostID(postID),
		Author:    user,
		Content:   content,
		Like:      0,
		Dislike:   0,
		CreatedAt: time.Now(),
	}
	err = h.service.CheckCommentInput(comm)
	if err != nil {
		errorHeader(w, "comment is invalid", http.StatusBadRequest)
		return
	}
	err = h.service.CreateComment(comm)
	if err != nil {
		errorHeader(w, "comment is not created", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, PostAddress+postIDStr, http.StatusSeeOther)
}
