package handler

import (
	"forumv2"
	"forumv2/internal/service"
	"log"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: &service,
	}
}

func (h *Handler) InitRoutes() {
	router := http.NewServeMux()

	router.HandleFunc("/", h.index)
	router.HandleFunc("/sign-in", h.userSignIn)
	router.HandleFunc("/sign-up", h.userSignUp)
	router.HandleFunc("/logout", h.IsAuthorized(h.logOutHandler))
	router.HandleFunc("/create-post", h.IsAuthorized(h.createPost))
	router.HandleFunc("/comments/", h.comment)
	router.HandleFunc("/like-comment", h.LikeComment)
	router.HandleFunc("/myprofile", h.myprofile)

	router.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("../internal/template/"))))
	// router.Handle("/template/img/", http.StripPrefix("/template/img/", http.FileServer(http.Dir("/template/img"))))
	srv := new(forumv2.Server)
	if err := srv.Run("8081", router); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
