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
		errorHeader(w, "", http.StatusMethodNotAllowed)
		//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	uuidCtx := r.Context().Value("uuid")
	if uuidCtx == nil {
		errorHeader(w, "", http.StatusBadRequest)
		//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	uuid := uuidCtx.(uuid.UUID)
	user, err := h.service.GetUsersInfoByUUIDService(uuid)
	if err != nil {
		errorHeader(w, "", http.StatusInternalServerError)
		//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	postIDStr := r.FormValue("postID")
	if postIDStr == "" {
		errorHeader(w, "", http.StatusBadRequest)
		//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	content := r.FormValue("content")
	if content == "" {
		errorHeader(w, "message is too long", http.StatusBadRequest)
		//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		errorHeader(w, "", http.StatusBadRequest)
		//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	t := time.Now()
	timeFormat := t.Format("15:04:04,02 January 2006")

	comm := models.Comment{
		PostID:    postID,
		Author:    user.Username,
		Content:   content,
		Like:      0,
		Dislike:   0,
		CreatedAt: timeFormat,
	}
	err = h.service.CheckCommentInput(comm)
	if err != nil {
		errorHeader(w, "comment is invalid", http.StatusBadRequest)
		//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = h.service.CreateCommentsInService(comm)
	if err != nil {
		errorHeader(w, "comment is not created", http.StatusInternalServerError)
		//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, PostAddress+postIDStr, http.StatusSeeOther)
}
