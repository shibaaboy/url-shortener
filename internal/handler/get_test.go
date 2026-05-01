package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/shibaaboy/url-shortener/internal/storage"
)

func TestGetHandler(t *testing.T) {
	store := storage.NewStorage()

	store.Set("abc", "https://example.com")

	h := NewHandler(store)

	r := chi.NewRouter()
	r.Get("/{id}", h.GetHandler)

	req := httptest.NewRequest(http.MethodGet, "/abc", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("expected 307, got %d", res.StatusCode)
	}

	if res.Header.Get("Location") != "https://example.com" {
		t.Errorf("wrong redirect location")
	}
}
