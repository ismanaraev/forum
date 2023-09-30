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

type SignlePost struct {
	Data     models.User
	Post     models.Post
	Comments []models.Comment
}

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
	user, _ := h.getUserFromRequest(r)
	res := SignlePost{
		Data:     user,
		Post:     post,
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
		user, err := h.getUserFromRequest(r)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusInternalServerError)
			return
		}
		data := AllData{
			Data:       user,
			Categories: categories,
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		user, err := h.getUserFromRequest(r)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusInternalServerError)
			return
		}
		err = r.ParseMultipartForm(models.MaxPictureSizeBytes)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		title, ok := r.MultipartForm.Value["title"]
		if !ok {
			log.Printf("title is empty")
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		content, ok := r.MultipartForm.Value["content"]
		if !ok {
			log.Print("content is empty")
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		categoriesArr := r.MultipartForm.Value["categories"]
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
		var pictureList []models.Picture
		if r.MultipartForm.File != nil {
			pictures := r.MultipartForm.File["picture"]
			if len(pictures) > models.MaxPicturesPerPost {
				errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			for i := range pictures {
				fileHeader := pictures[i]
				file, err := fileHeader.Open()
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
				pictureList = append(pictureList, models.Picture{Value: string(temp), Size: int(fileHeader.Size), Type: imageType})
			}
		}

		post := models.Post{
			Title:      title[0],
			Content:    content[0],
			Author:     user,
			CreatedAt:  time.Now(),
			Categories: categories,
			Pictures:   pictureList,
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

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		user, err := h.getUserFromRequest(r)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		postId := r.URL.Query().Get("post")
		postIdInt, err := strconv.ParseInt(postId, 10, 64)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		post, err := h.service.GetPostByIDinService(models.PostID(postIdInt))
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		if post.Author != user {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		tmpl, err := template.ParseFiles(TemplateDir+"html/editPost.html", TemplateDir+"html/header.html")
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		categories, err := h.service.GetAllCategories()
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		data := struct {
			Data       models.User
			Post       models.Post
			Categories []models.Category
		}{
			Data:       user,
			Post:       post,
			Categories: categories,
		}
		err = tmpl.Execute(w, &data)
		if err != nil {
			log.Print(err)
		}
	case http.MethodPost:
		user, err := h.getUserFromRequest(r)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		err = r.ParseForm()
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		postIdStr, ok := r.Form["postID"]
		if !ok {
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		title, ok := r.Form["title"]
		if !ok {
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		content, ok := r.Form["content"]
		if !ok {
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		postID, err := strconv.ParseInt(postIdStr[0], 10, 64)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		post, err := h.service.GetPostByIDinService(models.PostID(postID))
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		if post.Author != user {
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		post.Title = title[0]
		post.Content = content[0]
		err = h.service.UpdatePost(post)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, PostAddress+strconv.Itoa(int(postID)), http.StatusSeeOther)
	default:
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	user, err := h.getUserFromRequest(r)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	postIdStr := r.URL.Query().Get("post")
	postID, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	post, err := h.service.GetPostByIDinService(models.PostID(postID))
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if post.Author != user {
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = h.service.DeletePostByID(post.ID)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, IndexAddress, http.StatusSeeOther)
}
