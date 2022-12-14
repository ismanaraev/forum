package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"forum3/internal/models"
	"forum3/internal/repository"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

const TOKEN_SECRET = 15

var (
	UserNotCreated     = errors.New("User not created in Service")
	PasswordNotHashing = errors.New("Password is not hashing in Service")
	UUIDNotCreated     = errors.New("UUID is not created")
)

type AuthService struct {
	repo repository.Authorization
	gen  uuid.Generator
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		gen:  uuid.NewGen(),
		repo: repo,
	}
}

// Добавление нового юзера в базу,запрос на репу
func (a *AuthService) CreateUserService(user models.Auth) (int, error) {
	var err error

	if !checkValid(user) {
		return http.StatusBadRequest, fmt.Errorf("Create user in service: %w", UserNotCreated)
	}

	user.Uuid, err = a.gen.NewV4()
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Create user in service: %w", UUIDNotCreated)
	}

	user.Password, err = generateHashPassword(user.Password)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Create user in service: %w", PasswordNotHashing)
	}

	return a.repo.CreateUser(user)
}

// Проверка на авторизацию
func (a *AuthService) AuthorizationUserService(user models.Auth) (string, error) {
	var err error
	checkUser, err := a.repo.GetUserInfo(user)
	if err != nil {
		return "User is exist", err
	}

	if checkUser.Email != user.Email {
		return "Not correct emali", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(user.Password))
	if err != nil {
		return "Not correct password", err
	}

	value, err := a.CreateSession(checkUser)
	if err != nil {
		return "Session not created", err
	}
	return value, err
}

// Cоздает токен и время токена и отправляет в БД
func (a *AuthService) CreateSession(user models.Auth) (string, error) {
	token := CreateToken()
	expireTime := time.Now()
	return token, a.repo.SetSession(user, token, expireTime)
}

func CreateToken() string {
	b := make([]byte, TOKEN_SECRET)
	if _, err := rand.Read(b); err != nil {
		log.Fatalf("Token for user not created")
	}
	return hex.EncodeToString(b)
}
