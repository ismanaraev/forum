package repository

import (
	"database/sql"
	"forumv2/internal/models"
	"time"
)

type Repository struct {
	User
	Post
	Session
	Comments
	Reactions
}

type User interface {
	SetSession(user models.User, token string, time time.Time) error
	CreateUser(user models.User) (int, error)
	GetUserInfo(user models.User) (models.User, error)
	GetUsersEmail(user models.User) (models.User, error)
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
	GetPostsByCategory(category models.Category) ([]models.Post, error)
	FilterPostsByMultipleCategories(categories []models.Category) ([]models.Post, error)
	CreateCategory(name string) error
	GetCategoryByName(name string) (models.Category, error)
	UpdatePost(models.Post) error
}

type Session interface {
	GetSessionFromDB(token string) (models.UserID, error)
	DeleteSessionFromDB(models.UserID) error
}

type Comments interface {
	CreateComments(models.Comment) error

	GetAllComments() ([]models.Comment, error)
	GetCommentsByID(postID models.PostID) ([]models.Comment, error)
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

func NewRepository(db *sql.DB) Repository {
	return Repository{
		User:      NewUserSQLite(db),
		Post:      NewPostSQLite(db),
		Session:   NewSessionSQLite(db),
		Comments:  NewCommentsSQLite(db),
		Reactions: NewReactionsSQLite(db),
	}
}
