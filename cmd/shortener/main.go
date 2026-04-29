package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/shibaaboy/url-shortener/internal/config"
	"github.com/shibaaboy/url-shortener/internal/handler"
	"github.com/shibaaboy/url-shortener/internal/storage"
)

func main() {
	config.ParseFlags()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	srv := newApp()

	fmt.Println("Server address", config.Addr())
	fmt.Println("Base URL", config.BaseURL())

	go startServer(srv)
	return waitForShutdown(srv)
}

func startServer(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func newApp() *http.Server {
	store := storage.NewStorage()
	h := handler.NewHandler(store)
	r := initRouter(h)

	return &http.Server{
		Addr:    config.Addr(),
		Handler: r,
	}
}

func waitForShutdown(srv *http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}

func initRouter(h *handler.Handler) http.Handler {

	r := chi.NewRouter()
	r.Post("/", h.PostHandler)
	r.Get("/{id}", h.GetHandler)
	return r

}
