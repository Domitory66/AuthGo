package service

import "github.com/Domitory66/AuthGo/pkg/repository"

type Auth interface {
}

type Service struct {
	Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
