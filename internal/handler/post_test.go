package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/shibaaboy/url-shortener/internal/storage"
)

func TestPostHandler(t *testing.T) {
	store := storage.NewStorage()

	h := NewHandler(store)

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("https://example.com"))
	req.Host = "localhost:8080"

	w := httptest.NewRecorder()

	h.PostHandler(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected 201, got %d", resp.StatusCode)
	}
}
