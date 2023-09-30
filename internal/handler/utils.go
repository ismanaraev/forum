package handler

import (
	"errors"
	"fmt"
	"forumv2/internal/models"
	"log"
	"net/http"
)

func (h *Handler) getUserFromRequest(r *http.Request) (models.User, error) {
	idCtx := r.Context().Value(MiddlewareUID)
	if idCtx == nil {
		log.Printf("userID context is nil")
		return models.User{}, errors.New("no context value")
	}
	id := idCtx.(models.UserID)
	user, err := h.service.GetUsersInfoByUUIDService(id)
	if err != nil {
		log.Print(err)
		return models.User{}, fmt.Errorf("failed to get user from service: %v", err)
	}
	return user, nil
}
