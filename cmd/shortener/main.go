package main

import (
	"internal/app"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.RootRouter)
	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
