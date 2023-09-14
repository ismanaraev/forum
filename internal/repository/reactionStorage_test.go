package repository_test

import (
	db "forumv2/db/SQLite3"
	"forumv2/internal/models"
	"forumv2/internal/repository"
	"reflect"
	"testing"
	"time"

	"github.com/gofrs/uuid"
)

func TestCreateLikeForPost(t *testing.T) {
	data, err := db.Database("./test.db")
	defer RemoveDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := repository.NewRepository(data)
	id, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	user := models.User{
		ID:       models.UserID(id),
		Name:     "user",
		Username: "username",
		Email:    "asdf@asdf.com",
		Password: "asdfasdf",
	}
	err = repo.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}
	expectedPost := models.Post{
		ID:        models.PostID(1),
		Title:     "title",
		Content:   "content",
		Author:    user,
		Like:      1,
		CreatedAt: time.Now().Truncate(time.Second),
	}
	expectedPostId, err := repo.CreatePost(expectedPost)
	if err != nil {
		t.Fatal(err)
	}
	like := models.LikePost{
		UserID: user.ID,
		PostID: expectedPostId,
		Status: models.Like,
	}
	_, err = repo.CreateLikeForPost(like)
	if err != nil {
		t.Fatal(err)
	}
	stmt, err := data.Prepare(`SELECT userID,postID,status FROM likePost WHERE postID = $1`)
	if err != nil {
		t.Fatal(err)
	}
	row := stmt.QueryRow(expectedPostId)
	var gotLike models.LikePost
	var tempUsrIDStr string
	err = row.Scan(&tempUsrIDStr, &gotLike.PostID, &gotLike.Status)
	if err != nil {
		t.Fatal(err)
	}
	gotLike.UserID, err = models.UserIDFromString(tempUsrIDStr)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(gotLike, like) {
		t.Logf("like mismatch: got %v, want %v", gotLike, like)
		t.Fail()
	}
}

func TestCreateLikeForComment(t *testing.T) {
	data, err := db.Database("./test.db")
	defer RemoveDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := repository.NewRepository(data)
	id, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	user := models.User{
		ID:       models.UserID(id),
		Name:     "user",
		Username: "username",
		Email:    "asdf@asdf.com",
		Password: "asdfasdf",
	}
	err = repo.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}
	expectedPost := models.Post{
		ID:        models.PostID(1),
		Title:     "title",
		Content:   "content",
		Author:    user,
		CreatedAt: time.Now().Truncate(time.Second),
	}
	expectedPostId, err := repo.CreatePost(expectedPost)
	if err != nil {
		t.Fatal(err)
	}
	expectedComment := models.Comment{
		ID:        1,
		PostID:    expectedPostId,
		Author:    models.User{ID: user.ID},
		Content:   "content",
		Like:      1,
		CreatedAt: time.Now().Truncate(time.Second),
	}
	err = repo.CreateComment(expectedComment)
	if err != nil {
		t.Fatal(err)
	}
	like := models.LikeComment{
		UserID:     user.ID,
		CommentsID: expectedComment.ID,
		Status:     models.Like,
	}
	_, err = repo.CreateLikeForComment(like)
	if err != nil {
		t.Fatal(err)
	}
	stmt, err := data.Prepare(`SELECT userID,commentsID,status FROM likeComments WHERE commentsID = $1`)
	if err != nil {
		t.Fatal(err)
	}
	row := stmt.QueryRow(expectedPostId)
	var gotLike models.LikeComment
	var tempStr string
	err = row.Scan(&tempStr, &gotLike.CommentsID, &gotLike.Status)
	if err != nil {
		t.Fatal(err)
	}
	gotLike.UserID, err = models.UserIDFromString(tempStr)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(gotLike, like) {
		t.Logf("like mismatch: got %v, want %v", gotLike, like)
		t.Fail()
	}
}
