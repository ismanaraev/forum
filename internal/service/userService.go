package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"forumv2/internal/models"
	"log"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

const TOKEN_SECRET = 15

type userService struct {
	repo User
	gen  uuid.Generator
}

func newUserService(repo User) *userService {
	return &userService{
		gen:  uuid.NewGen(),
		repo: repo,
	}
}

// Добавление нового юзера в базу,запрос на репу
func (u *userService) CreateUserService(user models.User) error {
	var err error

	if !userValidation(user) {
		return fmt.Errorf("Create user in service: %w", err)
	}
	userUuid, err := u.gen.NewV4()
	if err != nil {
		return fmt.Errorf("Create user in service: %w", err)
	}
	user.ID = models.UserID(userUuid)
	user.Password, err = generateHashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("Create user in service: %w", err)
	}
	return u.repo.CreateUser(user)
}

// Проверка на авторизацию
func (u *userService) AuthorizationUserService(user models.User) (string, error) {
	var err error
	checkUser, err := u.repo.GetUserInfoByEmail(user.Email)
	if err != nil {
		return "User is exist", err
	}
	if checkUser.Email != user.Email {
		return "Not correct email", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(user.Password))
	if err != nil {
		return "Not correct password", err
	}

	value, err := u.CreateSessionService(checkUser)
	if err != nil {
		return "Session not created", err
	}
	return value, err
}

// Получения данных юзера из БД
func (u *userService) GetUserInfoService(user models.User) (models.User, error) {
	userInfo, err := u.repo.GetUserInfoByEmail(user.Email) // Получает информацию с помощью почты
	if err != nil {
		return models.User{}, err
	}
	return userInfo, nil
}

// Получения данных юзера из БД c помощью токена
func (u *userService) GetUsersInfoByUUIDService(id models.UserID) (models.User, error) {
	userInfo, err := u.repo.GetUsersInfoByUUID(id)
	if err != nil {
		return models.User{}, err
	}
	return userInfo, nil
}

// Cоздает токен и время токена и отправляет в БД
func (u *userService) CreateSessionService(user models.User) (string, error) {
	token := CreateToken()
	expireTime := time.Now()
	return token, u.repo.SetSession(user, token, expireTime)
}

func CreateToken() string {
	b := make([]byte, TOKEN_SECRET)
	if _, err := rand.Read(b); err != nil {
		log.Print("Token for user not created")
	}
	return hex.EncodeToString(b)
}

func (u *userService) CheckUserEmail(email string) (bool, error) {
	return u.repo.CheckUserEmail(email)
}

func (u *userService) CheckUserUsername(username string) (bool, error) {
	return u.repo.CheckUserUsername(username)
}
