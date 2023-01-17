package handler

import (
	"net/http"
	"text/template"
)

func (h *Handler) needToSign(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/need-to-sign" {
		errorHeader(w, "", http.StatusNotFound)
		//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		errorHeader(w, "", http.StatusMethodNotAllowed)
		//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	html, err := template.ParseFiles(TemplateDir + "html/needtoSign.html")
	if err != nil {
		errorHeader(w, "", http.StatusNotFound)
		//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	html.Execute(w, "")

}
