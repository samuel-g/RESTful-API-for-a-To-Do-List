package handlers

import (
	"encoding/json"
	"go-todo-app/models"
	"net/http"

	"github.com/gorilla/mux"
)

var todos []models.Todo

func GetTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			var todo models.Todo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todos = append(todos, todo)
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}
