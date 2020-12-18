package localdb

import (
	"time"

	"github.com/elemes1/gobridge_2/domain"
)

var Todos = []domain.Todo{{ID: 12, User: &domain.User{}, TaskName: "My First Task", Description: "First Task Description", Created: time.Now(), Updated: time.Now(), Completed: false}}

var user = domain.User{ID: 1, Name: "Samuel"}

func NewTodoStore() *TodoStore {
	return &TodoStore{
		todos: Todos,
	}
}

type TodoStore struct {
	todos []domain.Todo
}

func (s *TodoStore) List() ([]domain.Todo, error) {
	return s.todos, nil
}
