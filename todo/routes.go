package todo

import (
	"net/http"
)

type route struct {
	path    string
	handler http.HandlerFunc
}

var routes = []route{
	{
		path:    "/ping",
		handler: pong,
	},
}
