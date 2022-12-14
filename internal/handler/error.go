package handler

import (
	"html/template"
	"net/http"
)

type Response struct {
	NumberOfError int
	Message       string
}

func PrintErrors(w http.ResponseWriter, ms string, code int) {
	resp := Response{
		NumberOfError: code,
		Message:       ms,
	}

	html, err := template.ParseFiles("./ui/error.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error2", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	err = html.Execute(w, resp)
	if err != nil {
		http.Error(w, "500 Internal Server Error3", http.StatusInternalServerError)
		return
	}
}
