package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shibaaboy/url-shortener/internal/storage"
)

func TestAPIShortenHandler(t *testing.T) {
	store := storage.NewStorage()
	h := NewHandler(store)

	body := []byte(`{"url":"https://practicum.yandex.ru"}`)

	req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	h.APIShortenHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", res.StatusCode)
	}

	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Fatalf("expected application/json, got %s", ct)
	}

	var resp struct {
		Result string `json:"result"`
	}

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		t.Fatalf("decode error: %v", err)
	}

	if resp.Result == "" {
		t.Fatalf("expected result URL, got empty")
	}
}
