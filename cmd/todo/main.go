package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/elemes1/gobridge_2/api"
	"github.com/elemes1/gobridge_2/store/localdb"
)

func main() {
	store := localdb.NewTodoStore()
	r := api.NewRouter(api.Resources{
		Storage: store,
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8900",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting http Server ")
	err := srv.ListenAndServe()
	fmt.Println("program exit: ", err)
}
