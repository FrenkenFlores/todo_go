package repository

import (
	todogo "github.com/FrenkenFlores/todo_go"
	"github.com/jmoiron/sqlx"
)

type Authentication interface {
	CreateUser(user todogo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authentication
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authentication: NewAuthPostgres(db),
	}
}
