package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(rsc Resources) *mux.Router {
	r := mux.NewRouter()

	for _, rt := range getRoutes(rsc) {
		r.HandleFunc(rt.path, rt.handler).Methods(rt.method)
	}
	http.Handle("/", r)
	return r
}
