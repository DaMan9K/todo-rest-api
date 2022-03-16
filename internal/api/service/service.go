package service

import (
	"github.com/DaMan9K/todo-app/internal/api/repository"
	"github.com/DaMan9K/todo-app/internal/todo"
	"github.com/DaMan9K/todo-app/internal/user"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
