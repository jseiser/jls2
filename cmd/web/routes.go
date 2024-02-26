package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Health Check Routes
	mux.Handle("GET /api/live", app.live())
	mux.Handle("GET /api/ready", app.ready())
	// REST Routes
	mux.Handle("GET /api/todo", app.getAllTodos())
	mux.Handle("GET /api/todo/{id}", app.getTodoByID())
	mux.Handle("POST /api/todo", app.postTodo())
	mux.Handle("PUT /api/todo/{id}", app.putTodoByID())
	mux.Handle("DELETE /api/todo/{id}", app.deleteTodoByID())

	// Catch All Route
	mux.Handle("GET /", app.not_found())

	standard := alice.New(app.recoverPanic, app.httpLogger, secureHeaders)

	return standard.Then(mux)
}
