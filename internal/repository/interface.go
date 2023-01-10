package repository

import (
	"database/sql"
	"forumv2/internal/models"
	"time"

	"github.com/gofrs/uuid"
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
	GetUsersInfoByUUID(id uuid.UUID) (models.User, error) //++
}

type Post interface {
	GetAllPost() ([]models.Post, error)
	GetPostByID(id int) (models.Post, error)
	GetUsersPost(uuid uuid.UUID) ([]models.Post, error)
	GetPostWithCategory(category string) ([]models.Post, error)
	GetPostIdWithUUID(uuid uuid.UUID) ([]int, error)
	CreatePost(post models.Post) (int, error)
	GetUsersLikePosts(i []int) ([]models.Post, error)
}

type Session interface {
	GetSessionFromDB(token string) (uuid.UUID, error)
	DeleteSessionFromDB(uuid.UUID) error
}

type Comments interface {
	GetAllComments() ([]models.Comment, error)
	GetCommentsByID(postID int) ([]models.Comment, error)
	CreateComments(models.Comment) (int, error)
}

type Reactions interface {
	CreateLikeForPost(like models.LikePost) (models.LikePost, error)
	CreateLikeForComment(like models.LikeComment) (models.LikeComment, error)
	UpdatePostLikeStatus(like models.LikePost) (models.LikePost, error)
	UpdateCommentLikeStatus(like models.LikeComment) (models.LikeComment, error)
	GetUserIDfromLikePost(like models.LikePost) (int, error)
	GetLikeStatusByPostAndUserID(like models.LikePost) (models.LikeStatus, error)
	GetLikeStatusByCommentAndUserID(like models.LikeComment) (models.LikeStatus, error)
	IncrementPostLikeByPostID(postID int) error
	DecrementPostLikeByPostID(postID int) error
	IncrementPostDislikeByPostID(postID int) error
	DecrementPostDislikeByPostID(postID int) error
	IncrementCommentLikeByCommentsID(commentID int) error
	DecrementCommentLikeByCommentsID(commentID int) error
	DecrementCommentDislikeByCommentsID(commentID int) error
	IncrementCommentDislikeByCommentsID(commentID int) error
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
