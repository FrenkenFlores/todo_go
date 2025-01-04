package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
