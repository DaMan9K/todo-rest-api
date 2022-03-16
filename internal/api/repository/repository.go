package repository

import (
	"github.com/DaMan9K/todo-app/internal/todo"
	"github.com/DaMan9K/todo-app/internal/user"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GetUser(username, password string) (user.User, error)
}
type TodoList interface {
	Create(userID int, list todo.TodoList) (int, error)
	GetAll(userID int) ([]todo.TodoList, error)
	GetByID(userID, listId int) (todo.TodoList, error)
	Update(userId, listId int, input todo.UpdateListInput) error
	Delete(userID, listId int) error
}
type TodoItem interface {
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
