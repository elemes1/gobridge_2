package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elemes1/gobridge_2/domain"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestShowTodo_ListsAllTodosInTheDatabase(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		todosFromDB      []domain.Todo
		expectedTaskID   int
		expectedTaskName string
	}{
		{
			name: "finds task by ID",
			todosFromDB: []domain.Todo{
				{ID: 12, TaskName: "task"},
			},
			expectedTaskID:   12,
			expectedTaskName: "task",
		},
		{
			name: "finds task by ID when multiple exist",
			todosFromDB: []domain.Todo{
				{ID: 44, TaskName: "task44"},
				{ID: 12, TaskName: "task12"},
			},
			expectedTaskID:   12,
			expectedTaskName: "task12",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			store := StorageMock{
				ListFn: func() ([]domain.Todo, error) {
					return tt.todosFromDB, nil
				},
			}
			fn := showTodoHandler(store)

			w := httptest.NewRecorder()
			r, err := http.NewRequest("GET", "/", nil)
			require.Nil(t, err)
			r = mux.SetURLVars(r, map[string]string{
				"id": "12",
			})

			// act
			fn(w, r)

			// assert
			resp := w.Result()
			require.Equal(t, 200, resp.StatusCode)

			b, err := ioutil.ReadAll(resp.Body)
			require.Nil(t, err)
			defer resp.Body.Close()

			var todo domain.Todo
			err = json.Unmarshal(b, &todo)
			require.Nil(t, err)

			require.Equal(t, tt.expectedTaskID, todo.ID, "id does not match")
			require.Equal(t, tt.expectedTaskName, todo.TaskName, "the task name does not match")
		})
	}
}
