package service

import (
	authorization "github.com/Domitory66/AuthGo"
	"github.com/Domitory66/AuthGo/pkg/repository"
)

type Auth interface {
	CreateUser(user authorization.User) (int, error)
	GenerateToken(username, password string) (string, error)
	//ParseToken(token string) (int, error)
}

type Service struct {
	Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos.Authorization),
	}
}
