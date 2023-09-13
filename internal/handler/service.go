package handler

import "forumv2/internal/models"

type Service interface {
	User
	Post
	Session
	Comments
	Reactions
}

type User interface {
	CreateSessionService(user models.User) (string, error)
	CreateUserService(user models.User) error
	AuthorizationUserService(models.User) (string, error)
	GetUserInfoService(user models.User) (models.User, error)
	GetUsersInfoByUUIDService(id models.UserID) (models.User, error)
	CheckUserEmail(email string) (bool, error)
	CheckUserUsername(username string) (bool, error)
}

type Post interface {
	CreatePostService(post models.Post) (models.PostID, error)
	GetAllPostService() ([]models.Post, error)
	GetUsersPostInService(uuid models.UserID) ([]models.Post, error)
	GetUsersLikedPosts(id models.UserID) ([]models.Post, error)
	GetPostByIDinService(id models.PostID) (models.Post, error)
	FilterPostsByCategories([]string) ([]models.Post, error)
	CreateCategory(string) error
	GetCategoryByName(string) (models.Category, error)
	CheckPostInput(models.Post) error
}

type Session interface {
	DeleteSessionService(id models.UserID) error
	GetSessionService(token string) (models.UserID, error)
}

type Comments interface {
	GetCommentsByPostID(postID models.PostID) ([]models.Comment, error)
	CreateComment(com models.Comment) error
	CheckCommentInput(models.Comment) error
}

type Reactions interface {
	LikePostService(like models.LikePost) error
	LikeCommentService(like models.LikeComment) error
}
