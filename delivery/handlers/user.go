package handlers

import "net/http"

func (h Handler) AddUserHandler(w http.ResponseWriter, r *http.Request)   {}
func (h Handler) AuthorizeHandler(w http.ResponseWriter, r *http.Request) {}
func (h Handler) LogOutHandler(w http.ResponseWriter, r *http.Request)    {}
