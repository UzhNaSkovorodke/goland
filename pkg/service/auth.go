package service

import (
	todo "awesomeProject"
	"awesomeProject/pkg/repository"
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

const salt = "asddfhghkjvghhfdgdfg"
const signInKey = "stone_best"
const tokenTTL = 12 * time.Hour

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   strconv.Itoa(user.Id),
		},
		user.Id,
	})

	return token.SignedString([]byte(signInKey))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
