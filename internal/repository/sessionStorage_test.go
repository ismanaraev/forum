package repository_test

import (
	"database/sql"
	"errors"
	db "forumv2/db/SQLite3"
	"forumv2/internal/models"
	"forumv2/internal/repository"
	"testing"
	"time"

	"github.com/gofrs/uuid"
)

func TestDeleteSessionFromDB(t *testing.T) {
	data, err := db.Database(DB_PATH)
	defer RemoveDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := repository.NewRepository(data)
	uuid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	id := models.UserID(uuid)
	expectedUser := models.User{
		ID:       id,
		Name:     "name",
		Username: "username",
		Email:    "email@email.com",
		Password: "password",
	}
	err = repo.CreateUser(expectedUser)
	if err != nil {
		t.Fatal(err)
	}
	token := "123456"
	expire := time.Now().Add(time.Hour)
	err = repo.SetSession(expectedUser, token, expire)
	if err != nil {
		t.Fatal(err)
	}
	gotId, err := repo.GetSessionFromDB(token)
	if err != nil {
		t.Fatal(err)
	}
	if gotId.String() != expectedUser.ID.String() {
		t.Fail()
	}
	err = repo.DeleteSessionFromDB(expectedUser.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = repo.GetSessionFromDB(token)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			t.Fatal(err)
		}
	}
}
