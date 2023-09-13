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

func TestCreateComment(t *testing.T) {
	data, err := db.Database(DB_PATH)
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
		Name:     "name",
		Username: "username",
		Email:    "email@email.com",
		Password: "password",
	}
	err = repo.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}
	post := models.Post{
		ID:        models.PostID(1),
		Title:     "title",
		Content:   "content",
		Author:    models.User{ID: user.ID},
		CreatedAt: time.Now().Truncate(time.Second),
	}
	_, err = repo.CreatePost(post)
	if err != nil {
		t.Fatal(err)
	}
	comment := models.Comment{
		ID:        1,
		PostID:    models.PostID(1),
		Author:    models.User{ID: user.ID},
		Content:   "content",
		CreatedAt: time.Now().Truncate(time.Second),
	}
	err = repo.CreateComment(comment)
	if err != nil {
		t.Fatal(err)
	}
	gotComment, err := repo.GetCommentByCommentID(1)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(gotComment, comment) {
		t.Logf("comment mismatch: got %v want %v", gotComment, comment)
		t.Fail()
	}
}

func TestGetCommentsByPostID(t *testing.T) {
	data, err := db.Database(DB_PATH)
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
		Name:     "name",
		Username: "username",
		Email:    "email@email.com",
		Password: "password",
	}
	err = repo.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}
	post := models.Post{
		ID:        models.PostID(1),
		Title:     "title",
		Content:   "content",
		Author:    models.User{ID: user.ID},
		CreatedAt: time.Now().Truncate(time.Second),
	}
	_, err = repo.CreatePost(post)
	if err != nil {
		t.Fatal(err)
	}
	post2 := models.Post{
		ID:        models.PostID(2),
		Title:     "title2",
		Content:   "content2",
		Author:    models.User{ID: user.ID},
		CreatedAt: time.Now().Truncate(time.Second),
	}
	_, err = repo.CreatePost(post2)
	if err != nil {
		t.Fatal(err)
	}
	expectedComments := []models.Comment{
		{
			ID:        1,
			PostID:    models.PostID(1),
			Author:    models.User{ID: user.ID},
			Content:   "content",
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        2,
			PostID:    models.PostID(1),
			Author:    models.User{ID: user.ID},
			Content:   "content2",
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        3,
			PostID:    models.PostID(1),
			Author:    models.User{ID: user.ID},
			Content:   "content3",
			CreatedAt: time.Now().Truncate(time.Second),
		},
	}
	additionalPost := models.Comment{
		ID:        4,
		PostID:    models.PostID(2),
		Author:    models.User{ID: user.ID},
		Content:   "content4",
		CreatedAt: time.Now().Truncate(time.Second),
	}
	for _, val := range expectedComments {
		err = repo.CreateComment(val)
		if err != nil {
			t.Fatal(err)
		}
	}
	err = repo.CreateComment(additionalPost)
	if err != nil {
		t.Fatal(err)
	}
	gotComments, err := repo.GetCommentsByPostID(post.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(gotComments) != len(expectedComments) {
		t.Logf("got and expected length mismatch: got %v, want %v", len(gotComments), len(expectedComments))
		t.Fail()
	}
	for i := range gotComments {
		if !reflect.DeepEqual(gotComments[i], expectedComments[i]) {
			t.Logf("comment mismatch: got %v, want %v", gotComments[i], expectedComments[i])
			t.Fail()
		}
	}
}

func TestUpdateComment(t *testing.T) {
	data, err := db.Database(DB_PATH)
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
		Name:     "name",
		Username: "username",
		Email:    "email@email.com",
		Password: "password",
	}
	err = repo.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}
	post := models.Post{
		ID:        models.PostID(1),
		Title:     "title",
		Content:   "content",
		Author:    models.User{ID: user.ID},
		CreatedAt: time.Now().Truncate(time.Second),
	}
	_, err = repo.CreatePost(post)
	if err != nil {
		t.Fatal(err)
	}
	comment := models.Comment{
		ID:        1,
		PostID:    models.PostID(1),
		Author:    models.User{ID: user.ID},
		Content:   "content",
		CreatedAt: time.Now().Truncate(time.Second),
	}
	err = repo.CreateComment(comment)
	if err != nil {
		t.Fatal(err)
	}
	updatedComment := models.Comment{
		ID:        1,
		PostID:    models.PostID(1),
		Author:    models.User{ID: user.ID},
		Content:   "content3",
		CreatedAt: time.Now().Truncate(time.Second),
	}
	err = repo.UpdateComment(updatedComment)
	if err != nil {
		t.Fatal(err)
	}
	gotComment, err := repo.GetCommentByCommentID(1)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(gotComment, updatedComment) {
		t.Logf("comment mismatch: got %v want %v", gotComment, updatedComment)
		t.Fail()
	}
}
