package repository

import (
	"database/sql"
	"fmt"
	"forumv2/internal/models"
)

type reactionsStorage struct {
	db *sql.DB
}

func newReactionsSQLite(db *sql.DB) *reactionsStorage {
	return &reactionsStorage{
		db: db,
	}
}

func (r *reactionsStorage) CreateLikeForPost(like models.LikePost) (models.LikePost, error) {
	queryForLike, err := r.db.Prepare(`INSERT INTO likePost(userID,postID, status) VALUES ($1,$2,$3)`)
	if err != nil {
		return like, fmt.Errorf("[ReactionStorage]:Error with CreateLikeForPost method in repository: %w", err)
	}
	_, err = queryForLike.Exec(like.UserID.String(), like.PostID, like.Status)
	if err != nil {
		return like, fmt.Errorf("[ReactionStorage]:Error with CreateLikeForPost method in repository: %v", err)
	}
	return like, nil
}

func (r *reactionsStorage) CreateLikeForComment(like models.LikeComment) (models.LikeComment, error) {
	queryForLike, err := r.db.Prepare(`INSERT INTO likeComments(userID, commentsID, status) VALUES ($1,$2,$3)`)
	if err != nil {
		return like, fmt.Errorf("[ReactionStorage]:Error with CreateLikeForComment method in repository: %w", err)
	}
	_, err = queryForLike.Exec(like.UserID.String(), like.CommentsID, like.Status)
	if err != nil {
		return like, fmt.Errorf("[ReactionStorage]:Error with CreateLikeForComment method in repository: %v", err)
	}
	return like, nil
}

func (r *reactionsStorage) UpdatePostLikeStatus(like models.LikePost) error {
	records := ("UPDATE likePost SET status = $1 WHERE postID = $2")
	query, err := r.db.Prepare(records)
	if err != nil {
		return fmt.Errorf("[ReactionStorage]:Error with UpdatePostLikeStatus method in repository: %v", err)
	}
	_, err = query.Exec(like.Status, like.PostID)
	if err != nil {
		return fmt.Errorf("[ReactionStorage]:Error with UpdatePostLikeStatus method in repository: %v", err)
	}
	return nil
}

func (r *reactionsStorage) UpdateCommentLikeStatus(like models.LikeComment) error {
	records := ("UPDATE likeComments SET status = $1 WHERE commentsID = $2")
	query, err := r.db.Prepare(records)
	if err != nil {
		return fmt.Errorf("[ReactionStorage]:Error with UpdateCommentLikeStatus method in repository: %v", err)
	}
	_, err = query.Exec(like.Status, like.CommentsID)
	if err != nil {
		return fmt.Errorf("[ReactionStorage]:Error with UpdateCommentLikeStatus method in repository: %v", err)
	}
	return nil
}

func (r *reactionsStorage) GetUserIDfromLikePost(like models.LikePost) (models.PostID, error) {
	row := r.db.QueryRow("SELECT postID FROM likePost WHERE userID=$1", like.UserID)
	temp := models.LikePost{}
	err := row.Scan(&temp.PostID)
	if err != nil {
		return temp.PostID, fmt.Errorf("[ReactionStorage]:Error with GetUserIDfromLikePost method in repository: %v", err)
	}
	return temp.PostID, nil
}

func (r *reactionsStorage) GetLikeStatusByPostAndUserID(like models.LikePost) (models.LikeStatus, error) {
	stmt := `SELECT status FROM likePost WHERE userID == $1 AND postID == $2`
	query, err := r.db.Prepare(stmt)
	if err != nil {
		return models.NoLike, fmt.Errorf("[ReactionStorage]:Error with GetLikeStatusByPostAndUserID method in repository: %v", err)
	}
	res := query.QueryRow(like.UserID.String(), like.PostID)

	var status models.LikeStatus
	err = res.Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.NoLike, nil
		}
		return models.NoLike, fmt.Errorf("[ReactionStorage]:Error with GetLikeStatusByPostAndUserID method in repository: %v", err)
	}
	return status, nil
}

func (r *reactionsStorage) GetLikeStatusByCommentAndUserID(like models.LikeComment) (models.LikeStatus, error) {
	stmt := `SELECT status FROM likeComments WHERE userID == $1 AND commentsID == $2`
	query, err := r.db.Prepare(stmt)
	if err != nil {
		return models.NoLike, fmt.Errorf("[ReactionStorage]:Error with GetLikeStatusByCommentAndUserID method in repository: %v", err)
	}
	res := query.QueryRow(like.UserID.String(), like.CommentsID)
	var status models.LikeStatus
	err = res.Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.NoLike, nil
		}
		return models.NoLike, fmt.Errorf("[ReactionStorage]:Error with GetLikeStatusByCommentAndUserID method in repository: %v", err)
	}
	return status, nil
}

func (r *reactionsStorage) DeleteCommentLike(like models.LikeComment) error {
	stmt := `DELETE FROM likeComments WHERE commentsID == $1 AND userID == $2`
	query, err := r.db.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = query.Exec(like.CommentsID, like.UserID.String())
	if err != nil {
		return err
	}
	return nil
}

func (r *reactionsStorage) DeletePostLike(like models.LikePost) error {
	stmt := `DELETE FROM likePost WHERE postID == $1 AND userID == $2`
	query, err := r.db.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = query.Exec(like.PostID, like.UserID.String())
	if err != nil {
		return err
	}
	return nil
}
