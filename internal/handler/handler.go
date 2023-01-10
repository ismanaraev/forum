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

const (
	SignUpAddress        = "/sign-up"
	SignInAddress        = "/sign-in"
	LogoutAddress        = "/logout"
	CreatePostAddress    = "/create-post"
	CommentsAddress      = "/comments/"
	LikePostAddress      = "/like-post"
	LikeCommentAddress   = "/like-comment"
	CreateCommentAddress = "/create-comment"
	MyProfileAddress     = "/myprofile"
	PostAddress          = "/post/"
	TemplateAddress      = "/template/"
	TemplateDir          = "../internal/template/"
)

func (h *Handler) InitRoutes() {
	router := http.NewServeMux()

	router.HandleFunc("/", h.index)
	router.HandleFunc(SignInAddress, h.userSignIn)
	router.HandleFunc(SignUpAddress, h.userSignUp)
	router.HandleFunc(LogoutAddress, h.IsAuthorized(h.logOutHandler))
	router.HandleFunc(CreatePostAddress, h.IsAuthorized(h.CreatePost))
	router.HandleFunc(PostAddress, h.Post)
	router.HandleFunc(LikePostAddress, h.IsAuthorized(h.LikePost))
	router.HandleFunc(LikeCommentAddress, h.IsAuthorized(h.LikeComment))
	router.HandleFunc(CreateCommentAddress, h.IsAuthorized(h.CreateComment))
	router.HandleFunc(MyProfileAddress, h.myprofile)

	router.Handle(TemplateAddress, http.StripPrefix("/template/", http.FileServer(http.Dir(TemplateDir))))
	srv := new(forumv2.Server)
	if err := srv.Run("8081", router); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
