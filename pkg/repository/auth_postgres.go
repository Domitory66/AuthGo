package repository

import (
	"fmt"

	authorization "github.com/Domitory66/AuthGo"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user authorization.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) values ($1,$2) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username string) (authorization.User, error) {
	var user authorization.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1", userTable)
	err := r.db.Get(&user, query, username)
	if err != nil {
		return user, err
	}
	query = fmt.Sprintf("SELECT password_hash FROM %s WHERE username=$1", userTable)
	err = r.db.Get(&user.Password, query, username)
	return user, err
}
