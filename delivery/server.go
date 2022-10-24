package delivery

import (
	"forum/delivery/handlers"
	"forum/internal/service"
	"net/http"
)

func Serve(service service.Service) {
	handlers := handlers.NewHandler(service)
	mux := NewMux(handlers)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func NewMux(handler handlers.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", handler.IsAuthorized(handler.AddUserHandler))
	mux.HandleFunc("/authorize", handler.IsAuthorized(handler.AuthorizeHandler))
	mux.HandleFunc("/logout", handler.IsAuthorized(handler.LogOutHandler))
	return mux
}
