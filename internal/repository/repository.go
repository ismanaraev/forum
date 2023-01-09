package repository

import (
	"database/sql"
	"forum3/internal/models"
	"time"

	"github.com/gofrs/uuid"
)

type Repository struct {
	Authorization
	Post
	Session
}

type Authorization interface {
	SetSession(user models.Auth, token string, time time.Time) error
	CreateUser(user models.Auth) (int, error)
	GetUserInfo(user models.Auth) (models.Auth, error)
	GetUsersEmail(user models.Auth) (models.Auth, error)
	GetUsersInfoByUUID(id uuid.UUID) (models.Auth, error)
}

type Post interface {
	GetAllPost() ([]models.Post, error)
	GetAllComments() ([]models.Comments, error)
	GetPostByID(id int) (models.Post, error)
	GetCommentsByID(postID int) ([]models.Comments, error)
	CreateComments(models.Comments) (int, error)
	CreatePost(post models.Post) (int, error)
	UpdatePost(post models.Post) (int, error)
	DeletePost(post models.Post) (int, error)
	CreateLikeForPost(like models.LikePost) (models.LikePost, error)
	CreateLikeForComment(like models.LikeComments) (models.LikeComments, error)
	UpdatePostLikeStatus(like models.LikePost) (models.LikePost, error)
	UpdateCommentLikeStatus(like models.LikeComments) (models.LikeComments, error)
	GetUUIDbyUser(like models.LikePost) int
	GetLikeStatusByPostAndUserID(like models.LikePost) (models.LikeStatus, error)
	GetLikeStatusByCommentAndUserID(like models.LikeComments) (models.LikeStatus, error)
	// CounterLike() int
}

type Session interface {
	GetSessionFromDB(token string) (uuid.UUID, error)
	DeleteSessionFromDB(uuid.UUID) error
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		Authorization: NewAuthSQLite(db),
		Post:          NewPostSQLite(db),
		Session:       NewSessionSQLite(db),
	}
}
