package models

import (
	"time"
)

type Todo struct {
	ID       int       `json:"id,omitempty"`
	Content  string    `json:"content,omitempty"`
	Created  time.Time `json:"created,omitempty"`
	Complete bool      `json:"complete,omitempty"`
}

var DB = []Todo{}

type NewTodo struct {
	Content  string `json:"content,omitempty"`
	Complete bool   `json:"complete,omitempty"`
}

type UpdateTodo struct {
	Content  string `json:"content,omitempty"`
	Complete bool   `json:"complete,omitempty"`
}

type DeleteTodo struct {
	ID int `json:"id,omitempty"`
}

type GetTodo struct {
	ID int `json:"id,omitempty"`
}

func (t *Todo) Insert(newtodo NewTodo) (int, error) {

	createTodo := &Todo{
		ID:       len(DB) + 1,
		Content:  newtodo.Content,
		Created:  time.Now(),
		Complete: newtodo.Complete,
	}

	DB = append(DB, createTodo)

	return createTodo.ID, nil

}

// Need to get ALL, optionally with various filters
// func (t *Todo) GetAll(id int) (Todo, error) {
// 	return nil, nil
// }

// func (t *Todo) Get(gettodo GetTodo) (Todo, error) {
// 	return nil, nil
// }

// func (t *Todo) Delete(deletetodo DeleteTodo) (Todo, error) {
// 	return nil, nil
// }

// func (t *Todo) Update(updatetodo UpdateTodo) (Todo, error) {
// 	return nil, nil
// }
