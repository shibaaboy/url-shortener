package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	value, ok := h.store.Get(id)

	if ok {
		http.Redirect(w, r, value, http.StatusTemporaryRedirect)
	} else {
		http.Error(w, "Error. Short url not found", http.StatusBadRequest)
	}
}
