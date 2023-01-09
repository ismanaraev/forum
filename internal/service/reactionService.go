package service

import (
	"forumv2/internal/models"
	"forumv2/internal/repository"
)

type ReactionsService struct {
	repo repository.Reactions
}

func NewReactionsService(repo repository.Reactions) *ReactionsService {
	return &ReactionsService{
		repo: repo,
	}
}

func (r *ReactionsService) LikePostService(like models.LikePost) (models.LikePost, error) {
	status, err := r.repo.GetLikeStatusByPostAndUserID(like)
	if err != nil {
		return models.LikePost{}, err
	}
	if status == models.NoLike {
		if _, err := r.repo.CreateLikeForPost(like); err != nil {
			return models.LikePost{}, err
		}
		switch like.Status {
		case models.Like:
			r.repo.IncrementPostLikeByPostID(like.PostID)
		case models.DisLike:
			r.repo.IncrementPostDislikeByPostID(like.PostID)
		}
	} else {
		if status != like.Status {
			if _, err := r.repo.UpdatePostLikeStatus(like); err != nil {
				return models.LikePost{}, err
			}
			if err := r.addLikePost(like, status); err != nil {
				return models.LikePost{}, err
			}
		}
	}
	return like, nil
}

func (r *ReactionsService) LikeCommentService(like models.LikeComment) (models.LikeComment, error) {
	status, err := r.repo.GetLikeStatusByCommentAndUserID(like)
	if err != nil {
		return models.LikeComment{}, err
	}
	if status == models.NoLike {
		if _, err := r.repo.CreateLikeForComment(like); err != nil {
			return models.LikeComment{}, err
		}
		switch like.Status {
		case models.Like:
			r.repo.IncrementCommentLikeByCommentsID(like.CommentsID)
		case models.DisLike:
			r.repo.IncrementCommentDislikeByCommentsID(like.CommentsID)
		}
	} else {
		if status != like.Status {
			if _, err := r.repo.UpdateCommentLikeStatus(like); err != nil {
				return models.LikeComment{}, err
			}
			if err := r.addLikeComment(like, status); err != nil {
				return models.LikeComment{}, err
			}
		}
	}
	return like, nil
}

func (r *ReactionsService) addLikePost(like models.LikePost, prev models.LikeStatus) error {
	switch like.Status {
	case models.Like:
		err := r.repo.IncrementPostLikeByPostID(like.PostID)
		if err != nil {
			return err
		}
		return r.repo.DecrementPostDislikeByPostID(like.PostID)
	case models.DisLike:
		err := r.repo.DecrementPostLikeByPostID(like.PostID)
		if err != nil {
			return err
		}
		return r.repo.IncrementPostDislikeByPostID(like.PostID)
	}
	return nil
}

func (r *ReactionsService) addLikeComment(like models.LikeComment, prev models.LikeStatus) error {
	switch like.Status {
	case models.Like:
		err := r.repo.IncrementCommentLikeByCommentsID(like.CommentsID)
		if err != nil {
			return err
		}
		return r.repo.DecrementCommentDislikeByCommentsID(like.CommentsID)
	case models.DisLike:
		err := r.repo.DecrementCommentLikeByCommentsID(like.CommentsID)
		if err != nil {
			return err
		}
		return r.repo.IncrementCommentDislikeByCommentsID(like.CommentsID)
	}
	return nil
}
