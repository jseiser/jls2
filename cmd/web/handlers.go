package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	models "github.com/jseiser/jls2/internal"
)

func (app *application) live() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("I am Alive"))
	})
}

func (app *application) ready() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("I am Ready"))
	})
}

func (app *application) not_found() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})
}

func (app *application) getAllTodos() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Get All Todos"))
	})
}

func (app *application) getTodoByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id < 1 {
			app.notFound(w)
			return
		}
		gettodo := models.GetTodo{
			ID: id,
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Get Todo by id: %d", gettodo)
	})
}

func (app *application) postTodo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var t models.NewTodo

		err := json.NewDecoder(r.Body).Decode(&t)

		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Put Todo by id: %v", t)
	})
}

func (app *application) putTodoByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id < 1 {
			app.notFound(w)
			return
		}

		var t models.UpdateTodo

		err = json.NewDecoder(r.Body).Decode(&t)

		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Put Todo by id: %v", t)
	})
}

func (app *application) deleteTodoByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id < 1 {
			app.notFound(w)
			return
		}
		deletetodo := models.DeleteTodo{
			ID: id,
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Get Todo by id: %d", deletetodo)
	})
}
