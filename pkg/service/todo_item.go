package service

import (
	todo "github.com/RymarSergey/my_todo"
	"github.com/RymarSergey/my_todo/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

//
//

//GetById(userId int, listId int) (todo.TodoItem, error)
//Delete(userId int, listId int, itemId int) error
//Update(userId int, listId int, itemId int, item todo.UpdateItemInput) error
func (s *TodoItemService) Create(userId int, listId int, item todo.TodoItem) (int, error) {
	if _, err := s.listRepo.GetById(userId, listId); err != nil {
		return 0, nil
	}
	return s.repo.Create(listId, item)
}
func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}
