package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/elemes1/gobridge_2/domain"
	"github.com/elemes1/gobridge_2/store/localdb"
	"github.com/gorilla/mux"
)

type Storage interface {
	List() ([]domain.Todo, error)
	// create(content domain.Todo) ([]todo.Todo, error)
	// get(attrib string, val string) ([]domain.Todo, error)
	// update(id int32, content todo.Todo) ([]domain.Todo, error)
	// remove() error
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(localdb.Todos)
}

func showTodoHandler(store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestVal, _ := strconv.Atoi(mux.Vars(r)["id"])
		w.Header().Set("Content-Type", "application/json")

		tgtTodo, code, err := showTodo(store, requestVal)
		if err != nil {
			http.Error(w, err.Error(), code)
		}

		json.NewEncoder(w).Encode(tgtTodo)
	}
}

func showTodo(store Storage, todoID int) (domain.Todo, int, error) {
	todos, err := store.List()
	if err != nil {
		return domain.Todo{}, http.StatusInternalServerError, fmt.Errorf("failed to get todos: %w", err)
	}
	for _, todo := range todos {
		if todo.ID == todoID {
			return todo, http.StatusOK, nil
		}
	}
	return domain.Todo{}, http.StatusNotFound, fmt.Errorf("todo with id not found")
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo domain.Todo
	// Todo add random Number generator
	todo.ID = 13
	_ = json.NewDecoder(r.Body).Decode(&todo)
	localdb.Todos = append(localdb.Todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestVal, _ := strconv.Atoi(mux.Vars(r)["id"])
	for i, todo := range localdb.Todos {
		if todo.ID == requestVal {
			localdb.Todos = append(localdb.Todos[:i], localdb.Todos[i+1:]...)
			todo.ID = requestVal
			_ = json.NewDecoder(r.Body).Decode(&todo)
			localdb.Todos = append(localdb.Todos, todo)
			break
		}
	}
	json.NewEncoder(w).Encode(true)

}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestVal, _ := strconv.Atoi(mux.Vars(r)["id"])
	for i, todo := range localdb.Todos {
		if todo.ID == requestVal {
			localdb.Todos = append(localdb.Todos[:i], localdb.Todos[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(true)
}
