package todo

import (
	"net/http"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	for _, rt := range routes {
		r.HandleFunc(rt.path, rt.handler).Methods(rt.method)
	}
	http.Handle("/", r)
	return r
}
