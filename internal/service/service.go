package service

import (
	"student-lms/internal/domain"
	"student-lms/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
