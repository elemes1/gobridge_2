package api

import (
	"net/http"
)

type route struct {
	path    string
	handler http.HandlerFunc
	method  string
}

type Resources struct {
	Storage Storage
}

func getRoutes(rsc Resources) []route {
	return []route{
		{
			path:    "/api/todos",
			handler: getTodos,
			method:  "GET",
		}, {
			path:    "/api/todos/{id}",
			handler: showTodoHandler(rsc.Storage),
			method:  "GET",
		},
		{
			path:    "/api/todos",
			handler: createTodo,
			method:  "POST",
		},
		{
			path:    "/api/todos/{id}",
			handler: updateTodo,
			method:  "PUT",
		},
		{
			path:    "/api/todos/{id}",
			handler: deleteTodo,
			method:  "DELETE",
		},
	}
}
