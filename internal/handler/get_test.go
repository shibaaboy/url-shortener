package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestGetHandler_OK(t *testing.T) {
	// подменяем storage
	storage = map[string]string{
		"abc": "https://example.com",
	}

	r := chi.NewRouter()
	r.Get("/{id}", GetHandler)

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
