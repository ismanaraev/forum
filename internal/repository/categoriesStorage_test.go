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

//func TestGetCategoriesByPostID(t *testing.T) {
//	data, err := db.Database("./test.db")
//	//defer RemoveDB()
//	if err != nil {
//		t.Fatal(err)
//	}
//	repo := repository.NewPostSQLite(data)
//	userRepo := repository.NewUserSQLite(data)
//	catRepo := repository.NewCategoriesStorage(data)
//	id, err := uuid.NewV4()
//	if err != nil {
//		t.Fatal(err)
//	}
//	user := models.User{
//		ID: models.UserID(id),
//	}
//	_, err = userRepo.CreateUser(user)
//	if err != nil {
//		t.Fatal(err)
//	}
//	cooking := models.Category{
//		ID:   models.CategoryID(1),
//		Name: "Cooking",
//	}
//	music := models.Category{
//		ID:   models.CategoryID(2),
//		Name: "Music",
//	}
//	Puters := models.Category{
//		ID:   models.CategoryID(3),
//		Name: "Puters",
//	}
//	err = catRepo.CreateCategory(cooking.Name)
//	if err != nil {
//		t.Fatal(err)
//	}
//	err = catRepo.CreateCategory(music.Name)
//	if err != nil {
//		t.Fatal(err)
//	}
//	err = catRepo.CreateCategory(Puters.Name)
//	if err != nil {
//		t.Fatal(err)
//	}
//	expectedPosts := []models.Post{
//		{
//			ID:        models.PostID(1),
//			Title:     "title",
//			Content:   "content",
//			Author:    models.User{ID: user.ID},
//			CreatedAt: time.Now().Truncate(time.Second),
//		},
//		{
//			ID:        models.PostID(2),
//			Title:     "title2",
//			Content:   "content2",
//			Author:    models.User{ID: user.ID},
//			CreatedAt: time.Now().Truncate(time.Second),
//		},
//		{
//			ID:        models.PostID(3),
//			Title:     "title3",
//			Content:   "content3",
//			Author:    models.User{ID: user.ID},
//			CreatedAt: time.Now().Truncate(time.Second),
//		},
//		{
//			ID:        models.PostID(4),
//			Title:     "title4",
//			Content:   "content4",
//			Author:    models.User{ID: user.ID},
//			CreatedAt: time.Now().Truncate(time.Second),
//		},
//	}
//	for i := range expectedPosts {
//		_, err = repo.CreatePost(expectedPosts[i])
//		if err != nil {
//			t.Fatal(err)
//		}
//	}
//	err = catRepo.AddCategoryToPost(expectedPosts[0].ID, cooking.ID)
//	if err != nil {
//		t.Fatal(err)
//	}
//	err = catRepo.AddCategoryToPost(expectedPosts[0].ID, music.ID)
//	if err != nil {
//		t.Fatal(err)
//	}
//	err = catRepo.AddCategoryToPost(expectedPosts[0].ID, Puters.ID)
//	if err != nil {
//		t.Fatal(err)
//	}
//	err = catRepo.AddCategoryToPost(expectedPosts[1].ID, cooking.ID)
//	if err != nil {
//		t.Fatal(err)
//	}
//	err = catRepo.AddCategoryToPost(expectedPosts[2].ID, cooking.ID)
//	if err != nil {
//		t.Fatal(err)
//	}
//	expected := []models.Category{cooking, music, Puters}
//	got, err := catRepo.GetCategoriesByPostID(expectedPosts[0].ID)
//	if err != nil {
//		t.Fatal(err)
//	}
//	if !reflect.DeepEqual(got, expected) {
//		t.Logf("categories array mismatch, want %v, got %v", got, expected)
//		t.Fail()
//	}
//}

func TestFilterByMultipleCategories(t *testing.T) {
	data, err := db.Database("./test.db")
	defer RemoveDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := repository.NewPostSQLite(data)
	userRepo := repository.NewUserSQLite(data)
	catRepo := repository.NewCategoriesStorage(data)
	id, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	user := models.User{
		ID: models.UserID(id),
	}
	_, err = userRepo.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}
	cooking := models.Category{
		ID:   models.CategoryID(1),
		Name: "Cooking",
	}
	music := models.Category{
		ID:   models.CategoryID(2),
		Name: "Music",
	}
	Puters := models.Category{
		ID:   models.CategoryID(3),
		Name: "Puters",
	}
	err = catRepo.CreateCategory(cooking.Name)
	if err != nil {
		t.Fatal(err)
	}
	err = catRepo.CreateCategory(music.Name)
	if err != nil {
		t.Fatal(err)
	}
	err = catRepo.CreateCategory(Puters.Name)
	if err != nil {
		t.Fatal(err)
	}
	expectedPosts := []models.Post{
		{
			ID:        models.PostID(1),
			Title:     "title",
			Content:   "content",
			Author:    models.User{ID: user.ID},
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(2),
			Title:     "title2",
			Content:   "content2",
			Author:    models.User{ID: user.ID},
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(3),
			Title:     "title3",
			Content:   "content3",
			Author:    models.User{ID: user.ID},
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(4),
			Title:     "title4",
			Content:   "content4",
			Author:    models.User{ID: user.ID},
			CreatedAt: time.Now().Truncate(time.Second),
		},
	}
	for i := range expectedPosts {
		_, err = repo.CreatePost(expectedPosts[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	err = catRepo.AddCategoryToPost(expectedPosts[0].ID, cooking.ID)
	if err != nil {
		t.Fatal(err)
	}
	err = catRepo.AddCategoryToPost(expectedPosts[0].ID, music.ID)
	if err != nil {
		t.Fatal(err)
	}
	err = catRepo.AddCategoryToPost(expectedPosts[0].ID, Puters.ID)
	if err != nil {
		t.Fatal(err)
	}
	err = catRepo.AddCategoryToPost(expectedPosts[1].ID, cooking.ID)
	if err != nil {
		t.Fatal(err)
	}
	err = catRepo.AddCategoryToPost(expectedPosts[1].ID, Puters.ID)
	if err != nil {
		t.Fatal(err)
	}
	err = catRepo.AddCategoryToPost(expectedPosts[2].ID, cooking.ID)
	if err != nil {
		t.Fatal(err)
	}
	expectedCookingAndPuters := []models.Post{expectedPosts[0], expectedPosts[1]}
	gotCookingAndPuters, err := repo.FilterPostsByMultipleCategories([]models.Category{cooking, Puters})
	if err != nil {
		t.Fatal(err)
	}
	if len(expectedCookingAndPuters) != len(gotCookingAndPuters) {
		t.Logf("expected and got length mismatch")
		t.Fail()
	}
	if !reflect.DeepEqual(expectedCookingAndPuters, gotCookingAndPuters) {
		t.Logf("expected and got post list mismatch: got %v, expected %v", gotCookingAndPuters, expectedCookingAndPuters)
		t.Fail()
	}
	expectedCookingMusicPuters := []models.Post{expectedPosts[0]}
	gotCookingMusicPuters, err := repo.FilterPostsByMultipleCategories([]models.Category{cooking, music, Puters})
	if err != nil {
		t.Fatal(err)
	}
	if len(expectedCookingMusicPuters) != len(gotCookingMusicPuters) {
		t.Logf("expected and got length mismatch")
		t.Fail()
	}
	if !reflect.DeepEqual(expectedCookingMusicPuters, gotCookingMusicPuters) {
		t.Logf("expected and got post list mismatch: got %v, expected %v", gotCookingMusicPuters, expectedCookingMusicPuters)
		t.Fail()
	}
}
