package service

import "github.com/FrenkenFlores/todo_go/pkg/repository"

type Authentication interface {
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
	return &Service{}
}
