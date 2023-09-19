package handler

import (
	"database/sql"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"forumv2/internal/models"
)

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	postid := strings.TrimPrefix(r.URL.Path, PostAddress)
	postID, err := strconv.ParseInt(postid, 10, 64)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	post, err := h.service.GetPostByIDinService(models.PostID(postID))
	if err != nil {
		log.Print(err)
		if errors.Is(err, sql.ErrNoRows) {
			errorHeader(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	comments, err := h.service.GetCommentsByPostID(post.ID)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles(TemplateDir+"html/commentPage.html", TemplateDir+"html/header.html")
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	idCtx := r.Context().Value(MiddlewareUID)
	if idCtx == nil {
		res := AllData{Data: models.User{}, Post: []models.Post{post}, Comments: comments}
		err = tmpl.Execute(w, res)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		return
	}
	id := idCtx.(models.UserID)
	user, err := h.service.GetUsersInfoByUUIDService(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	res := AllData{
		Data:     user,
		Post:     []models.Post{post},
		Comments: comments,
	}
	err = tmpl.Execute(w, res)
	if err != nil {
		log.Print(err)
	}
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles(TemplateDir+"html/createPost.html", TemplateDir+"html/header.html")
		if err != nil {
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		categories, err := h.service.GetAllCategories()
		if err != nil {
			log.Printf("failed to get all categories: %v", err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		data := AllData{
			Categories: categories,
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		idCtx := r.Context().Value(MiddlewareUID)
		if idCtx == nil {
			log.Printf("userID context is nil")
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		id := idCtx.(models.UserID)
		user, err := h.service.GetUsersInfoByUUIDService(id)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err = r.ParseMultipartForm(models.MaxPictureSizeBytes)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		title, ok := r.Form["title"]
		if !ok {
			log.Printf("title is empty")
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		content, ok := r.Form["content"]
		if !ok {
			log.Print("content is empty")
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		categoriesArr := r.Form["categories"]
		var categories []models.Category
		for _, val := range categoriesArr {
			cat, err := h.service.GetCategoryByName(val)
			if err != nil {
				log.Print(err)
				errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			categories = append(categories, cat)
		}
		file, fileHeader, err := r.FormFile("picture")
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		if fileHeader.Size > models.MaxPictureSizeBytes {
			log.Printf("file size is %v, expected less than %v", fileHeader.Size, models.MaxPictureSizeBytes)
			errorHeader(w, "Image size is too big, please try uploading smaller images", http.StatusBadRequest)
			return
		}
		temp := make([]byte, fileHeader.Size)
		n, err := file.Read(temp)
		if err != nil || n != int(fileHeader.Size) {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		imageType, err := models.StringToImageType(http.DetectContentType(temp))
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		post := models.Post{
			Title:      title[0],
			Content:    content[0],
			Author:     user,
			CreatedAt:  time.Now(),
			Categories: categories,
			Pictures:   models.Picture{Value: string(temp), Size: int(fileHeader.Size), Type: imageType},
			Like:       0,
			Dislike:    0,
		}
		err = h.service.CheckPostInput(post)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		postId, err := h.service.CreatePostService(post)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, PostAddress+strconv.FormatInt(int64(postId), 10), http.StatusSeeOther)
	default:
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
