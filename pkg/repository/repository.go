package repository

import (
	authorization "github.com/Domitory66/AuthGo"
	"github.com/jmoiron/sqlx"
)

const (
	userTable = "users"
)

type Authorization interface {
	CreateUser(authorization.User) (int, error)
	GetUser(username string) (authorization.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
