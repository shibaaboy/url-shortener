package handler

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/shibaaboy/url-shortener/internal/config"
	"github.com/shibaaboy/url-shortener/internal/model"
)

func (h *Handler) APIShortenHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 2048)
	defer r.Body.Close()

	var req model.ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}

	b := make([]byte, 6)
	_, _ = rand.Read(b)
	id := base64.URLEncoding.EncodeToString(b)

	h.store.Set(id, req.URL)

	shortURL := config.BaseURL() + "/" + id

	resp := model.ShortenResponse{Result: shortURL}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
