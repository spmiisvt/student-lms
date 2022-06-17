package service

import (
	"crypto/sha1"
	"fmt"
	"student-lms/internal/domain"
	"student-lms/internal/repository"
	"time"
)

const (
	salt       = "lbnwieubqibs923fhwe"
	signingKey = "ofweFGWQFFf23@323#w23qei"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
