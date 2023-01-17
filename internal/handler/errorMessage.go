package handler

import (
	"net/http"
	"text/template"
)

type ErrorBody struct {
	Status  int
	Message string
}

func errorHeader(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	errH := setError(status)
	html, err := template.ParseFiles(TemplateDir + "html/error.html")
	if err != nil {
		errorHeader(w, http.StatusInternalServerError)
		return
	}
	html.Execute(w, errH)
}

func setError(status int) *ErrorBody {
	return &ErrorBody{
		Status:  status,
		Message: http.StatusText(status),
	}
}
