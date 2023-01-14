package handler

import (
	"forumv2/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	uuidCtx := r.Context().Value("uuid")
	if uuidCtx == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	uuid := uuidCtx.(uuid.UUID)
	user, err := h.service.GetUsersInfoByUUIDService(uuid)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	postIDStr := r.FormValue("postID")
	if postIDStr == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	content := r.FormValue("content")
	if content == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	comm := models.Comment{
		PostID:    postID,
		Author:    user.Username,
		Content:   content,
		Like:      0,
		Dislike:   0,
		CreatedAt: time.Now().String(),
	}
	err = h.service.CreateCommentsInService(comm)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, PostAddress+postIDStr, http.StatusSeeOther)
}