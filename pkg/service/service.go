package service

import (
	todogo "github.com/FrenkenFlores/todo_go"
	"github.com/FrenkenFlores/todo_go/pkg/repository"
)

type Authentication interface {
	CreateUser(user todogo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authentication
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authentication: NewAuthService(repos.Authentication),
	}
}
