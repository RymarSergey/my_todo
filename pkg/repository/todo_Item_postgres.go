package repository

import (
	"fmt"
	"strings"

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
	getAllItemQuery := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on ti.id=li.item_id
													INNER JOIN %s ul on ul.list_id=li.list_id WHERE ul.user_id=$1 AND li.list_id=$2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	err := r.db.Select(&result, getAllItemQuery, userId, listId)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *TodoItemPostgres) GetById(userId int, itemId int) (todo.TodoItem, error) {
	result := todo.TodoItem{}

	getAllItemQuery := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on ti.id=li.item_id
													INNER JOIN %s ul on ul.list_id=li.list_id WHERE ul.user_id=$1 AND li.item_id=$2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	err := r.db.Get(&result, getAllItemQuery, userId, itemId)

	return result, err
}

func (r *TodoItemPostgres) Delete(userId int, itemId int) error {
	getAllItemQuery := fmt.Sprintf(`DELETE  FROM %s ti USING %s li, %s ul  WHERE ti.id=li.item_id AND ul.list_id=li.list_id AND ul.user_id=$1 AND li.item_id=$2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	_, err := r.db.Exec(getAllItemQuery, userId, itemId)

	return err
}

func (r *TodoItemPostgres) Update(userId int, itemId int, input todo.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argsId))
		args = append(args, *input.Title)
		argsId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argsId))
		args = append(args, *input.Description)
		argsId++
	}
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argsId))
		args = append(args, *input.Done)
		argsId++
	}

	setQuery := strings.Join(setValues, " ,")
	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id=li.item_id AND li.list_id=ul.list_id AND ul.user_id=$%d AND ti.id=$%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argsId, argsId+1)
	args = append(args, userId, itemId)

	fmt.Printf("Query: %s \n", query)
	fmt.Printf("Args: %s \n", args)

	result, err := r.db.Exec(query, args...)
	fmt.Printf("result: %+v \n", result)
	fmt.Printf("err: %s \n", err)

	return err
}
