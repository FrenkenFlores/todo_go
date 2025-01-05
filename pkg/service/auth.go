package service

import (
	"crypto/sha1"
	"fmt"

	todogo "github.com/FrenkenFlores/todo_go"
	"github.com/FrenkenFlores/todo_go/pkg/repository"
)

type AuthService struct {
	repo repository.Authentication
}

func NewAuthService(repo repository.Authentication) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todogo.User) (int, error) {
	user.Password = s.generatePassowrdHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePassowrdHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte("jiopqjgiop29edlkwjqio")))
}
