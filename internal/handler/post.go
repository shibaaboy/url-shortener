package handler

import (
	"io"
	"net/http"

	"github.com/shibaaboy/url-shortener/internal/config"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	originalURL := string(bodyBytes)
	id := "EwHXdJfB"
	storage[id] = originalURL

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(config.BaseURL() + r.Host + "/" + id))
}
