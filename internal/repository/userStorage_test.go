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

func TestCreateUser(t *testing.T) {
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
	stmt, err := data.Prepare(`SELECT ID,name,username,email,password FROM users WHERE ID = $1`)
	if err != nil {
		t.Fatal(err)
	}
	row := stmt.QueryRow(id.String())
	var gotUser models.User
	var tempStr string
	err = row.Scan(&tempStr, &gotUser.Name, &gotUser.Username, &gotUser.Email, &gotUser.Password)
	if err != nil {
		t.Fatal(err)
	}
	gotUser.ID, err = models.UserIDFromString(tempStr)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(gotUser, user) {
		t.Logf("user mismatch: got %v, want %v", gotUser, user)
		t.Fail()
	}
}

func TestGetUserInfoByEmail(t *testing.T) {
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
	expectedUser := models.User{
		ID:       models.UserID(id),
		Name:     "name",
		Username: "username",
		Email:    "email@email.com",
		Password: "password",
	}
	err = repo.CreateUser(expectedUser)
	if err != nil {
		t.Fatal(err)
	}
	id2, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	user2 := models.User{
		ID:       models.UserID(id2),
		Name:     "name",
		Username: "username2",
		Email:    "email2@email.com",
		Password: "password",
	}
	err = repo.CreateUser(user2)
	if err != nil {
		t.Fatal(err)
	}
	gotUser, err := repo.GetUserInfoByEmail("email@email.com")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(gotUser, expectedUser) {
		t.Logf("user mismatch: got %v want %v", gotUser, expectedUser)
		t.Fail()
	}
}

func TestGetUsersInfoByUUID(t *testing.T) {
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
	gotUser, err := repo.GetUsersInfoByUUID(id)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(gotUser, expectedUser) {
		t.Logf("user mismatch: got %v want %v", gotUser, expectedUser)
		t.Fail()
	}
}

func TestCheckUserEmail(t *testing.T) {
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
	got, err := repo.CheckUserEmail(expectedUser.Email)
	if err != nil {
		t.Fatal(err)
	}
	if !got {
		t.Fail()
	}
	got, err = repo.CheckUserEmail("totally@random.email")
	if err != nil {
		t.Fatal(err)
	}
	if got {
		t.Fail()
	}
}

func TestCheckUserUsername(t *testing.T) {
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
	got, err := repo.CheckUserUsername(expectedUser.Username)
	if err != nil {
		t.Fatal(err)
	}
	if !got {
		t.Fail()
	}
	got, err = repo.CheckUserUsername("totallyrandomusername")
	if err != nil {
		t.Fatal(err)
	}
	if got {
		t.Fail()
	}
}

func TestSetSession(t *testing.T) {
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
}
