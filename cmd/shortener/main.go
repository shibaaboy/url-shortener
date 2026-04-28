package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shibaaboy/url-shortener/internal/config"
	"github.com/shibaaboy/url-shortener/internal/handler"
)

func main() {
	config.ParseFlags()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	fmt.Println("Server address", config.Addr())
	fmt.Println("Base URL", config.BaseURL())

	r := chi.NewRouter()
	r.Post("/", handler.PostHandler)
	r.Get("/{id}", handler.GetHandler)

	return http.ListenAndServe(config.Addr(), r)
}
