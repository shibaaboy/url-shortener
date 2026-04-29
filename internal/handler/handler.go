package handler

import "github.com/shibaaboy/url-shortener/internal/storage"

type Handler struct {
	store *storage.Storage
}

func NewHandler(store *storage.Storage) *Handler {

	return &Handler{
		store: store,
	}
}
