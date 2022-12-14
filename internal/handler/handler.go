package handler

import (
	"forum3"
	"forum3/internal/service"
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

	router.HandleFunc("/home", h.homepage)
	router.HandleFunc("/sign-in", h.userSignIn)
	router.HandleFunc("/sign-up", h.userSignUp)
	router.HandleFunc("/logout", h.IsAuthorized(h.logOutHandler))
	router.HandleFunc("/post", h.IsAuthorized(h.postPage))
	router.HandleFunc("/create-post", h.createPost)
	router.HandleFunc("/update-post", h.updatePost)
	router.HandleFunc("/delete-post", h.deletePost)

	router.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("../internal/template/"))))
	router.Handle("/template/img/", http.StripPrefix("/template/img/", http.FileServer(http.Dir("/template/img"))))

	srv := new(forum3.Server)

	if err := srv.Run("8086", router); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}
