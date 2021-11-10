package todo

import "errors"

type TodoList struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}
type UserList struct {
	ID     int
	UserID int
	ListID int
}
type TodoItem struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title"  db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}
type UserItem struct {
	ID     int
	ListID int
	ItemID int
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        bool    `json:"done"`
}
type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (u *UpdateListInput) Validate() error {
	if u.Description == nil && u.Title == nil {
		return errors.New("update struct has no value")
	}
	return nil
}
