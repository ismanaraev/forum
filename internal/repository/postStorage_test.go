package repository_test

import (
	db "forumv2/db/SQLite3"
	"forumv2/internal/models"
	"forumv2/internal/repository"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/gofrs/uuid"
)

const DB_PATH = "./test.db"

func RemoveDB() {
	os.Remove(DB_PATH)
}

func TestCreatePost(t *testing.T) {
	data, err := db.Database("./test.db")
	defer RemoveDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := repository.NewPostSQLite(data)
	userRepo := repository.NewUserSQLite(data)
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
	_, err = userRepo.CreateUser(user)
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
	row := data.QueryRow(`SELECT ID, title, content, author, createdat, like, dislike FROM post WHERE ID = $1`, expectedPostId)
	var post models.Post
	var authorStr string
	var outTime int64
	err = row.Scan(&post.ID, &post.Title, &post.Content, &authorStr, &outTime, &post.Like, &post.Dislike)
	if err != nil {
		t.Fatal(err)
	}
	post.CreatedAt = time.Unix(outTime, 0)
	uuid, err := models.UserIDFromString(authorStr)
	if err != nil {
		t.Fatal(err)
	}
	gotUser, err := userRepo.GetUsersInfoByUUID(uuid)
	if err != nil {
		t.Fatal(err)
	}
	post.Author = gotUser
	if !reflect.DeepEqual(post, expectedPost) {
		t.Logf("CreatePost error, posts mismatch, got %v, want %v", post, expectedPost)
		t.Fail()
	}
}

func TestGetAllPosts(t *testing.T) {
	data, err := db.Database("./test.db")
	defer RemoveDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := repository.NewPostSQLite(data)
	userRepo := repository.NewUserSQLite(data)
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
	expectedPosts := []models.Post{
		{
			ID:        models.PostID(1),
			Title:     "title",
			Content:   "content",
			Author:    user,
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(2),
			Title:     "title2",
			Content:   "content2",
			Author:    user,
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(3),
			Title:     "title3",
			Content:   "content3",
			Author:    user,
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(4),
			Title:     "title4",
			Content:   "content4",
			Author:    user,
			CreatedAt: time.Now().Truncate(time.Second),
		},
	}
	for i := range expectedPosts {
		_, err = repo.CreatePost(expectedPosts[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	gotPosts, err := repo.GetAllPost()
	if err != nil {
		t.Fatal(err)
	}
	for i := range gotPosts {
		if !reflect.DeepEqual(gotPosts[i], expectedPosts[i]) {
			t.Logf("Post %v mismatch: got %v, want %v", i, gotPosts[i], expectedPosts[i])
			t.Fail()
		}
	}
}

func TestGetPostByID(t *testing.T) {
	data, err := db.Database("./test.db")
	defer RemoveDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := repository.NewPostSQLite(data)
	userRepo := repository.NewUserSQLite(data)
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
	expectedPosts := []models.Post{
		{
			ID:        models.PostID(1),
			Title:     "title",
			Content:   "content",
			Author:    user,
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(2),
			Title:     "title2",
			Content:   "content2",
			Author:    user,
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(3),
			Title:     "title3",
			Content:   "content3",
			Author:    user,
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(4),
			Title:     "title4",
			Content:   "content4",
			Author:    user,
			CreatedAt: time.Now().Truncate(time.Second),
		},
	}
	for i := range expectedPosts {
		_, err = repo.CreatePost(expectedPosts[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	for i := range expectedPosts {
		gotPost, err := repo.GetPostByID(expectedPosts[i].ID)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(gotPost, expectedPosts[i]) {
			t.Logf("Post mismatch, got %v, want %v", gotPost, expectedPosts[i])
			t.Fail()
		}
	}
}

func TestGetPostsByUserID(t *testing.T) {
	data, err := db.Database("./test.db")
	defer RemoveDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := repository.NewPostSQLite(data)
	userRepo := repository.NewUserSQLite(data)
	id, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	userA := models.User{
		ID: models.UserID(id),
	}
	_, err = userRepo.CreateUser(userA)
	if err != nil {
		t.Fatal(err)
	}
	userAPosts := []models.Post{
		{
			ID:        models.PostID(1),
			Title:     "title",
			Content:   "content",
			Author:    models.User{ID: userA.ID},
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(2),
			Title:     "title2",
			Content:   "content2",
			Author:    models.User{ID: userA.ID},
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(3),
			Title:     "title3",
			Content:   "content3",
			Author:    models.User{ID: userA.ID},
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(4),
			Title:     "title4",
			Content:   "content4",
			Author:    models.User{ID: userA.ID},
			CreatedAt: time.Now().Truncate(time.Second),
		},
	}
	for i := range userAPosts {
		_, err = repo.CreatePost(userAPosts[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	id2, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	userB := models.User{
		ID:       models.UserID(id2),
		Name:     "test",
		Username: "user",
		Email:    "email",
	}
	_, err = userRepo.CreateUser(userB)
	if err != nil {
		t.Fatal(err)
	}
	userBPosts := []models.Post{
		{
			ID:        models.PostID(5),
			Title:     "title5",
			Content:   "content5",
			Author:    models.User{ID: userB.ID},
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(6),
			Title:     "title6",
			Content:   "content6",
			Author:    models.User{ID: userB.ID},
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(7),
			Title:     "title7",
			Content:   "content7",
			Author:    models.User{ID: userB.ID},
			CreatedAt: time.Now().Truncate(time.Second),
		},
	}
	for i := range userBPosts {
		_, err = repo.CreatePost(userBPosts[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	gotUserAPosts, err := repo.GetPostsByUserID(userA.ID)
	if err != nil {
		t.Fatal(err)
	}
	for i := range gotUserAPosts {
		if !reflect.DeepEqual(gotUserAPosts[i], userAPosts[i]) {
			t.Logf("post mismatch, got %v, want %v", gotUserAPosts[i], userAPosts[i])
			t.Fail()
		}
	}
	gotUserBPosts, err := repo.GetPostsByUserID(userB.ID)
	if err != nil {
		t.Fatal(err)
	}
	for i := range gotUserBPosts {
		if !reflect.DeepEqual(gotUserBPosts[i], userBPosts[i]) {
			t.Logf("post mismatch, got %v, want %v", gotUserBPosts[i], userBPosts[i])
			t.Fail()
		}
	}
}

func TestGetUserLikedPosts(t *testing.T) {
	data, err := db.Database("./test.db")
	defer RemoveDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := repository.NewPostSQLite(data)
	userRepo := repository.NewUserSQLite(data)
	reactionRepo := repository.NewReactionsSQLite(data)
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
	expectedPosts := []models.Post{
		{
			ID:        models.PostID(1),
			Title:     "title",
			Content:   "content",
			Author:    user,
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(2),
			Title:     "title2",
			Content:   "content2",
			Author:    user,
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(3),
			Title:     "title3",
			Content:   "content3",
			Author:    user,
			CreatedAt: time.Now().Truncate(time.Second),
		},
		{
			ID:        models.PostID(4),
			Title:     "title4",
			Content:   "content4",
			Author:    user,
			CreatedAt: time.Now().Truncate(time.Second),
		},
	}
	for i := range expectedPosts {
		_, err = repo.CreatePost(expectedPosts[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	likes := []models.LikePost{
		{
			UserID: user.ID,
			PostID: models.PostID(1),
			Status: models.Like,
		},
		{
			UserID: user.ID,
			PostID: models.PostID(2),
			Status: models.DisLike,
		},
		{
			UserID: user.ID,
			PostID: models.PostID(3),
			Status: models.Like,
		},
	}
	expected := []models.Post{expectedPosts[0], expectedPosts[2]}
	for i := range likes {
		_, err = reactionRepo.CreateLikeForPost(likes[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	got, err := repo.GetUsersLikePosts(user.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != len(expected) {
		t.Fatalf("invalid length of liked post, got %v, want %v", len(got), len(expected))
	}
	for i := range got {
		if !reflect.DeepEqual(got[i], expected[i]) {
			t.Logf("post mismatch: got %v, want %v", got[i], expected[i])
			t.Fail()
		}
	}
}

func TestUpdatePost(t *testing.T) {
	data, err := db.Database("./test.db")
	defer RemoveDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := repository.NewPostSQLite(data)
	userRepo := repository.NewUserSQLite(data)
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
	_, err = userRepo.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}
	expectedPost := models.Post{
		ID:        models.PostID(1),
		Title:     "title",
		Content:   "content",
		Author:    models.User{ID: user.ID},
		CreatedAt: time.Now().Truncate(time.Second),
	}
	_, err = repo.CreatePost(expectedPost)
	if err != nil {
		t.Fatal(err)
	}
	expected := models.Post{
		ID:        models.PostID(1),
		Title:     "New title",
		Content:   "New Content",
		Author:    models.User{ID: user.ID},
		CreatedAt: time.Now().Truncate(time.Second),
	}
	err = repo.UpdatePost(expected)
	if err != nil {
		t.Fatal(err)
	}
	got, err := repo.GetPostByID(expected.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, expected) {
		t.Logf("post mismatch: got %v, want %v", got, expected)
		t.Fail()
	}
}

func TestGetPostsByCategory(t *testing.T) {
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
	second := models.Category{
		ID:   models.CategoryID(2),
		Name: "Second",
	}
	err = catRepo.CreateCategory("Cooking")
	if err != nil {
		t.Fatal(err)
	}
	err = catRepo.CreateCategory("Second")
	if err != nil {
		t.Fatal(err)
	}
	expectedCookingPosts := []models.Post{
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
	}
	expectedSecondPosts := []models.Post{
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
	for _, val := range expectedCookingPosts {
		_, err = repo.CreatePost(val)
		if err != nil {
			t.Fatal(err)
		}
		err = catRepo.AddCategoryToPost(val.ID, cooking.ID)
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, val := range expectedSecondPosts {
		_, err = repo.CreatePost(val)
		if err != nil {
			t.Fatal(err)
		}
		err = catRepo.AddCategoryToPost(val.ID, second.ID)
		if err != nil {
			t.Fatal(err)
		}
	}
	cookingPosts, err := catRepo.GetPostsByCategory(cooking)
	if err != nil {
		t.Fatal(err)
	}
	if len(cookingPosts) != len(expectedCookingPosts) {
		t.Fatalf("len of cookingPosts and expectedCookingPosts mismatch: got %v, want %v", len(cookingPosts), len(expectedCookingPosts))
	}
	for i := range cookingPosts {
		if !reflect.DeepEqual(cookingPosts[i], expectedCookingPosts[i]) {
			t.Logf("post mismatch, got %v, want %v", cookingPosts[i], expectedCookingPosts[i])
			t.Fail()
		}
	}
	secondPosts, err := catRepo.GetPostsByCategory(second)
	if err != nil {
		t.Fatal(err)
	}
	if len(secondPosts) != len(expectedSecondPosts) {
		t.Fatalf("len of secondPosts and expectedCookingPosts mismatch: got %v, want %v", len(secondPosts), len(expectedSecondPosts))
	}
	for i := range secondPosts {
		if !reflect.DeepEqual(secondPosts[i], expectedSecondPosts[i]) {
			t.Logf("post mismatch, got %v, want %v", secondPosts[i], expectedSecondPosts[i])
			t.Fail()
		}
	}
}
