package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/elemes1/gobridge_2/todo"
)

func main() {
	r := todo.NewRouter()

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
