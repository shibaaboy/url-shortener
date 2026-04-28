package handler

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/shibaaboy/url-shortener/internal/config"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 2048)
	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "request too large", http.StatusRequestEntityTooLarge)
		return
	}

	originalURL := string(bodyBytes)

	b := make([]byte, 6)
	_, _ = rand.Read(b)
	id := base64.URLEncoding.EncodeToString(b)

	storage[id] = originalURL

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte(config.BaseURL() + "/" + id))
}
