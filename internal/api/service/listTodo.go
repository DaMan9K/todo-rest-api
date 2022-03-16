package service

import (
	"github.com/DaMan9K/todo-app/internal/api/repository"
	"github.com/DaMan9K/todo-app/internal/todo"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userID int, list todo.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

func (s *TodoListService) GetAll(userID int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userID)
}
func (s *TodoListService) GetByID(userID, listId int) (todo.TodoList, error) {
	return s.repo.GetByID(userID, listId)
}
func (s *TodoListService) Update(userID, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userID, listId, input)
}
func (s *TodoListService) Delete(userID, listId int) error {
	return s.repo.Delete(userID, listId)
}
