package handler

import (
	"log"
	"net/http"

	"forumv2"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

const (
	IndexAddress          = "/"
	SignUpAddress         = "/sign-up"
	SignInAddress         = "/sign-in"
	LogoutAddress         = "/logout"
	CreatePostAddress     = "/create-post"
	LikePostAddress       = "/like-post"
	LikeCommentAddress    = "/like-comment"
	CreateCommentAddress  = "/create-comment"
	UpdatePostAddress     = "/update-post"
	DeletePostAddress     = "/delete-post"
	MyProfileAddress      = "/myprofile"
	PostAddress           = "/post/"
	TemplateAddress       = "/template/"
	TemplateDir           = "./template/"
	FilterAddress         = "/filter"
	SignatureCheck        = "/need-to-sign"
	CreateCategoryAddress = "/create-category"
	DeleteCategory        = "/delete-category"
)

func (h *Handler) InitRoutes(serverHost, serverPort string) {
	router := http.NewServeMux()

	router.HandleFunc("/", h.MayBeAuthorized(h.index))
	router.HandleFunc(SignInAddress, h.userSignIn)
	router.HandleFunc(SignUpAddress, h.userSignUp)
	router.HandleFunc(LogoutAddress, h.OnlyIfAuthorized(h.logOutHandler))
	router.HandleFunc(CreatePostAddress, h.OnlyIfAuthorized(h.CreatePost))
	router.HandleFunc(PostAddress, h.MayBeAuthorized(h.Post))
	router.HandleFunc(LikePostAddress, h.OnlyIfAuthorized(h.LikePost))
	router.HandleFunc(LikeCommentAddress, h.OnlyIfAuthorized(h.LikeComment))
	router.HandleFunc(CreateCommentAddress, h.OnlyIfAuthorized(h.CreateComment))
	router.HandleFunc(MyProfileAddress, h.OnlyIfAuthorized(h.myprofile))
	router.HandleFunc(FilterAddress, h.MayBeAuthorized(h.FilterByCategory))
	router.HandleFunc(SignatureCheck, h.needToSign)
	router.HandleFunc(CreateCategoryAddress, h.OnlyIfAuthorized(h.CreateCategory))
	router.HandleFunc(DeleteCategory, h.OnlyIfAuthorized(h.DeleteCategory))
	router.HandleFunc(UpdatePostAddress, h.OnlyIfAuthorized(h.UpdatePost))
	router.HandleFunc(DeletePostAddress, h.OnlyIfAuthorized(h.DeletePost))

	router.Handle(TemplateAddress, http.StripPrefix("/template/", http.FileServer(http.Dir(TemplateDir))))
	srv := new(forumv2.Server)
	if err := srv.Run(serverHost, serverPort, router); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
