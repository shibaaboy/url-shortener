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
	"github.com/shibaaboy/url-shortener/internal/logger"
	"github.com/shibaaboy/url-shortener/internal/middleware"
	"github.com/shibaaboy/url-shortener/internal/storage"
)

func main() {
	config.ParseFlags()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	store := storage.NewStorage()
	store.LoadFromFile(config.FileStoragePath())

	fmt.Println("Server address", config.Addr())
	fmt.Println("Base URL", config.BaseURL())
	logger.Initialize(config.LogLevel())

	srv := newApp(store)

	go startServer(srv)
	return waitForShutdown(srv, store)
}

func startServer(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func newApp(store *storage.Storage) *http.Server {
	h := handler.NewHandler(store)
	r := initRouter(h)

	return &http.Server{
		Addr:              config.Addr(),
		Handler:           r,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
}

func waitForShutdown(srv *http.Server, store *storage.Storage) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return err
	}

	store.SaveToFile(config.FileStoragePath())

	return nil
}

func initRouter(h *handler.Handler) http.Handler {
	r := chi.NewRouter()
	r.Use(logger.RequestLogger)
	r.Use(middleware.GzipMiddleware)

	r.Post("/", h.PostHandler)
	r.Post("/api/shorten", h.APIShortenHandler)
	r.Get("/{id}", h.GetHandler)

	return r
}
