package handler

import (
	"fmt"
	"forum3/internal/models"
	"html/template"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
)

// // авторизация
func (h *Handler) userSignIn(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sign-in" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	switch r.Method {

	case http.MethodGet:
		html, err := template.ParseFiles("../internal/template/html/signIn.html")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		html.Execute(w, nil)

	case http.MethodPost:

		r.ParseForm()

		email, ok := r.Form["email"]
		if !ok {
			http.Error(w, "username field not found", http.StatusInternalServerError)
			return
		}

		password, ok := r.Form["password"]
		if !ok {
			http.Error(w, "email field not found", http.StatusInternalServerError)
			return
		}

		data := models.Auth{
			Email:    email[0],
			Password: password[0],
		}

		token, err := h.service.AuthorizationUserService(data)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		cookie := http.Cookie{
			Name:   "session_name",
			Value:  token,
			MaxAge: 300 * 50,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/homepage", http.StatusSeeOther)

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

// регистрация
func (h *Handler) userSignUp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sign-up" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		html, err := template.ParseFiles("../internal/template/html/signUp.html")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		html.Execute(w, nil)

	case http.MethodPost:

		r.ParseForm()

		username, ok := r.Form["username"]
		if !ok {
			http.Error(w, "username field not found", http.StatusInternalServerError)
			return
		}

		name, ok := r.Form["name"]
		if !ok {
			http.Error(w, "name field not found", http.StatusInternalServerError)
			return
		}

		password, ok := r.Form["password"]
		if !ok {
			http.Error(w, "username field not found", http.StatusInternalServerError)
			return
		}

		email, ok := r.Form["email"]
		if !ok {
			http.Error(w, "email field not found", http.StatusInternalServerError)
			return
		}

		data := models.Auth{
			Name:     name[0],
			Username: username[0],
			Password: password[0],
			Email:    email[0],
		}

		status, err := h.service.Authorization.CreateUserService(data)
		if err != nil {
			http.Error(w, http.StatusText(status), status)
			log.Printf("User not created")
		}

		http.Redirect(w, r, "/sign-in", http.StatusSeeOther)

	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}

func (h *Handler) logOutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	uuidString := r.Context().Value("uuid")
	value := uuidString.(string)

	uuid, err := uuid.FromString(value)
	if err != nil {
		return
	}

	cookie := http.Cookie{
		Name:   "session_name",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)

	h.service.DeleteSessionRQtoRepo(uuid)
	http.Redirect(w, r, "/homepage", http.StatusSeeOther)
}
