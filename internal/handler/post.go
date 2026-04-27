package handler

import (
	"io"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		bodyBytes, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}

		originalURL := string(bodyBytes)
		id := "EwHXdJfB"
		storage[id] = originalURL

		w.WriteHeader(201)
		w.Write([]byte("https://" + r.Host + "/" + id))
	} else {
		http.Error(w, "Error", http.StatusMethodNotAllowed)
	}
}
