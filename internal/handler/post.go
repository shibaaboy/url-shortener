package handler

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/shibaaboy/url-shortener/internal/config"
)

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 2048)
	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error. Request too large", http.StatusBadRequest)
		return
	}

	originalURL := string(bodyBytes)

	b := make([]byte, 6)
	_, _ = rand.Read(b)
	id := base64.URLEncoding.EncodeToString(b)

	h.store.Set(id, originalURL)

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte(config.BaseURL() + "/" + id))
}
