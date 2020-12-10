package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/swtch1/gobridge_2/todo"
)

func main() {
	r := todo.NewRouter()

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()
	fmt.Println("program exit: ", err)
}
