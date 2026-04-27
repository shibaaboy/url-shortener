package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	value, ok := storage[id]

	if ok {
		http.Redirect(w, r, value, 307)
	} else {
		http.Error(w, "Error", http.StatusBadRequest)
	}
}
