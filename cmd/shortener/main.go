package main

import (
	"net/http"

	"github.com/shibaaboy/url-shortener/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.MainHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
