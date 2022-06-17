package repository

import (
	"github.com/jmoiron/sqlx"
	"student-lms/internal/domain"
	"student-lms/internal/repository/postgres"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
	}
}
