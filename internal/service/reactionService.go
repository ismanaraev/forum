package service

import (
	"forumv2/internal/models"
)

type reactionsService struct {
	repo        Reactions
	postRepo    Post
	commentRepo Comments
}

func newReactionsService(repo Reactions, postRepo Post, commentRepo Comments) *reactionsService {
	return &reactionsService{
		repo:        repo,
		postRepo:    postRepo,
		commentRepo: commentRepo,
	}
}

func (r *reactionsService) LikePostService(like models.LikePost) error {
	status, err := r.repo.GetLikeStatusByPostAndUserID(like)
	if err != nil {
		return err
	}
	post, err := r.postRepo.GetPostByID(like.PostID)
	if err != nil {
		return err
	}
	if status == models.NoLike {
		_, err := r.repo.CreateLikeForPost(like)
		if err != nil {
			return err
		}
		switch like.Status {
		case models.Like:
			post.Like += 1
		case models.DisLike:
			post.Dislike += 1
		}
		return r.postRepo.UpdatePost(post)
	}
	if status == like.Status {
		switch like.Status {
		case models.Like:
			post.Like -= 1
		case models.DisLike:
			post.Dislike -= 1
		}
		if err := r.repo.DeletePostLike(like); err != nil {
			return err
		}
		return r.postRepo.UpdatePost(post)
	}
	if status != like.Status {
		switch like.Status {
		case models.Like:
			post.Like += 1
			post.Dislike -= 1
		case models.DisLike:
			post.Like -= 1
			post.Dislike += 1
		}
		if err := r.postRepo.UpdatePost(post); err != nil {
			return err
		}
		return r.repo.UpdatePostLikeStatus(like)
	}
	return nil
}

func (r *reactionsService) LikeCommentService(like models.LikeComment) error {
	status, err := r.repo.GetLikeStatusByCommentAndUserID(like)
	if err != nil {
		return err
	}
	comment, err := r.commentRepo.GetCommentByCommentID(like.CommentsID)
	if err != nil {
		return err
	}
	if status == models.NoLike {
		if _, err := r.repo.CreateLikeForComment(like); err != nil {
			return err
		}
		switch like.Status {
		case models.Like:
			comment.Like += 1
		case models.DisLike:
			comment.Dislike += 1
		}
		err := r.commentRepo.UpdateComment(comment)
		if err != nil {
			return err
		}
		return nil
	}

	if status == like.Status {
		switch like.Status {
		case models.Like:
			comment.Like -= 1
		case models.DisLike:
			comment.Dislike -= 1
		}
		if err := r.commentRepo.UpdateComment(comment); err != nil {
			return err
		}
		if err := r.repo.DeleteCommentLike(like); err != nil {
			return err
		}
	}
	if status != like.Status {
		switch like.Status {
		case models.Like:
			comment.Like += 1
			comment.Dislike -= 1
		case models.DisLike:
			comment.Like -= 1
			comment.Dislike += 1
		}
		if err := r.commentRepo.UpdateComment(comment); err != nil {
			return err
		}
		return r.repo.UpdateCommentLikeStatus(like)
	}

	return nil
}
