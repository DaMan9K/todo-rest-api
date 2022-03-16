package todo

import "errors"

type TodoList struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}
type UsersList struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	ListID int `json:"list_id"`
}
type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title" `
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
type ListsItem struct {
	ID     int
	ListID int
	ItemID int
}
type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

//update structure has no values. args = nil

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
