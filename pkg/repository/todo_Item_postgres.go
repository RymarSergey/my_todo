package repository

import (
	"fmt"

	todo "github.com/RymarSergey/my_todo"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title,description,done) VALUES ($1,$2,$3) RETURNING id", todoItemsTable)
	row := r.db.QueryRow(createItemQuery, item.Title, item.Description, item.Done)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createListItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemQuery, listId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (r *TodoItemPostgres) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	result := []todo.TodoItem{}
	getAllItemQuery := fmt.Sprintf(`SELECT * FROM %s ti INNER JOIN %s li on ti.id=li.item_id
													INNER JOIN %s ul on ul.list_id=li.list_id WHERE ul.user_id=$1 AND li.list_id=$2`, todoItemsTable, listsItemsTable, usersListsTable)

	err := r.db.Select(&result, getAllItemQuery, userId, listId)
	if err != nil {
		return nil, err
	}
	return result, err
}
