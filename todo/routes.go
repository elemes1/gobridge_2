package todo

import (
	"net/http"
)

type route struct {
	path    string
	handler http.HandlerFunc
	method  string
}

var routes = []route{
	{
		path:    "/api/todos",
		handler: getTodos,
		method:  "GET",
	}, {
		path:    "/api/todos/{id}",
		handler: showTodo,
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
