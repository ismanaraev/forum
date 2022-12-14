package handler

import (
	"net/http"
)

func (h *Handler) postPage(w http.ResponseWriter, r *http.Request) {
	// email := r.Context().Value("email")
}

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {}
