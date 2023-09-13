package service

import (
	"forumv2/internal/models"
	"time"
)

type service struct {
	*userService
	*postService
	*sessionService
	*commentService
	*reactionsService
}

type Repository interface {
	User
	Post
	Session
	Comments
	Reactions
	Categories
}

func NewService(repo Repository) *service {
	return &service{
		newUserService(repo),
		newPostService(repo, repo, repo),
		newSessionService(repo),
		newCommentsService(repo),
		newReactionsService(repo, repo, repo),
	}
}

type User interface {
	SetSession(user models.User, token string, time time.Time) error
	CreateUser(user models.User) error
	GetUserInfoByEmail(email string) (models.User, error)
	GetUsersInfoByUUID(id models.UserID) (models.User, error) //++
	CheckUserEmail(email string) (bool, error)
	CheckUserUsername(username string) (bool, error)
}

type Post interface {
	CreatePost(post models.Post) (models.PostID, error)
	GetAllPost() ([]models.Post, error)
	GetPostByID(id models.PostID) (models.Post, error)
	GetPostsByUserID(uuid models.UserID) ([]models.Post, error)
	GetUsersLikePosts(id models.UserID) ([]models.Post, error)
	FilterPostsByMultipleCategories(categories []models.Category) ([]models.Post, error)
	UpdatePost(models.Post) error
}

type Session interface {
	GetSessionFromDB(token string) (models.UserID, error)
	DeleteSessionFromDB(models.UserID) error
}

type Comments interface {
	CreateComment(models.Comment) error
	GetCommentsByPostID(postID models.PostID) ([]models.Comment, error)
	GetCommentByCommentID(commentID int) (models.Comment, error)
	UpdateComment(models.Comment) error
}

type Reactions interface {
	CreateLikeForPost(like models.LikePost) (models.LikePost, error)
	CreateLikeForComment(like models.LikeComment) (models.LikeComment, error)

	GetUserIDfromLikePost(like models.LikePost) (models.PostID, error)
	GetLikeStatusByPostAndUserID(like models.LikePost) (models.LikeStatus, error)
	GetLikeStatusByCommentAndUserID(like models.LikeComment) (models.LikeStatus, error)

	UpdatePostLikeStatus(like models.LikePost) error
	UpdateCommentLikeStatus(like models.LikeComment) error

	DeletePostLike(models.LikePost) error
	DeleteCommentLike(models.LikeComment) error
}

type Categories interface {
	CreateCategory(string) error
	AddCategoryToPost(models.PostID, models.CategoryID) error
	GetCategoryByName(string) (models.Category, error)
	GetCategoriesByPostID(models.PostID) ([]models.Category, error)
	GetPostsByCategory(category models.Category) ([]models.Post, error)
}
