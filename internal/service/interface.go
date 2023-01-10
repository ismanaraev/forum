package service

import (
	"forumv2/internal/models"
	"forumv2/internal/repository"

	"github.com/gofrs/uuid"
)

type Service struct {
	User
	Post
	Session
	Comments
	Reactions
}

type User interface {
	CreateSessionService(user models.User) (string, error)
	CreateUserService(user models.User) (int, error)
	AuthorizationUserService(models.User) (string, error)
	GetUserInfoService(user models.User) (models.User, error)
	GetUsersInfoByUUIDService(id uuid.UUID) (models.User, error)
}

type Post interface {
	CreatePostService(post models.Post) (int, error)
	GetAllPostService(category string) ([]models.Post, error)
	GetUsersPostInService(uuid uuid.UUID) ([]models.Post, error)
	GetUserLikePostsInService(uuid uuid.UUID) ([]models.Post, error)
	GetPostByIDinService(id int) (models.Post, error)
}

type Session interface {
	DeleteSessionService(uuid uuid.UUID) error
	GetSessionService(token string) (uuid.UUID, error)
}

type Comments interface {
	GetAllCommentsInService() ([]models.Comment, error)
	GetCommentsByIDinService(postID int) ([]models.Comment, error)
	CreateCommentsInService(com models.Comment) (int, error)
}

type Reactions interface {
	LikePostService(like models.LikePost) (models.LikePost, error)
	LikeCommentService(like models.LikeComment) (models.LikeComment, error)
}

func NewService(repo repository.Repository) Service {
	return Service{
		User:      NewUserService(repo.User),
		Post:      NewPostService(repo.Post),
		Session:   NewSessionService(repo.Session),
		Comments:  NewCommentsService(repo.Comments),
		Reactions: NewReactionsService(repo.Reactions),
	}
}
