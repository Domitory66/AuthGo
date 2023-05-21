package service

import (
	"fmt"
	"time"

	authorization "github.com/Domitory66/AuthGo"
	"github.com/Domitory66/AuthGo/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const (
	salt       = "xdfvb4e5fytvyudu5"
	signingKey = "drtfd65udd5jg7"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user authorization.User) (int, error) {
	user.Password = s.generatePasswoordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username)
	if err != nil {
		return "", err
	}

	if err = s.ValidPaswoordHash(password, user.Password); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) generatePasswoordHash(password string) string {
	cost := 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	fmt.Println(string(hash))
	if err != nil {
		logrus.Errorf("Error generating pass")
	}
	return string(hash)
}

func (s *AuthService) ValidPaswoordHash(password_in, password_hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password_in))
	if err != nil {
		logrus.Errorf("Error password")
	}
	return err
}
