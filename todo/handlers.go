package todo

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//storage interface

type Storage interface {
	list() ([]Todo, error)
	create(content Todo) ([]Todo, error)
	get(attrib string, val string) ([]Todo, error)
	update(id int32, content Todo) ([]Todo, error)
	remove() error
}

//Todo Model
type Todo struct {
	ID          int       `json:"id"`
	User        *User     `json:"user"`
	TaskName    string    `json:"name"`
	Description string    `json:"description"`
	Created     time.Time `json:"creation_date"`
	Updated     time.Time `json:"completion_date"`
	Completed   bool      `json:"status"`
}

//Todo User
type User struct {
	ID   int    `json:"id"`
	Name string `json:"username"`
}

// sample Todo and User

var user = User{ID: 1, Name: "Samuel"}

var todos = []Todo{{ID: 12, User: &user, TaskName: "My First Task", Description: "First Task Description", Created: time.Now(), Updated: time.Now(), Completed: false}}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func showTodo(w http.ResponseWriter, r *http.Request) {
	requestVal, _ := strconv.Atoi(mux.Vars(r)["id"])
	w.Header().Set("Content-Type", "application/json")
	for _, todo := range todos {
		if todo.ID == requestVal {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	// Todo add random Number generator
	todo.ID = 13
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestVal, _ := strconv.Atoi(mux.Vars(r)["id"])
	for i, todo := range todos {
		if todo.ID == requestVal {
			todos = append(todos[:i], todos[i+1:]...)
			todo.ID = requestVal
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todos = append(todos, todo)
			break
		}
	}
	json.NewEncoder(w).Encode(true)

}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestVal, _ := strconv.Atoi(mux.Vars(r)["id"])
	for i, todo := range todos {
		if todo.ID == requestVal {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(true)
}
