package repository

import "github.com/jmoiron/sqlx"

type Authentication interface {
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
	return &Repository{}
}
