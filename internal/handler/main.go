package handler

import (
	"io"
	"net/http"
)

var storage = make(map[string]string)

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

func GetHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		currId := r.URL.Path
		currIdWithoutSlash := currId[1:]
		value, ok := storage[currIdWithoutSlash]

		if ok == true {
			http.Redirect(w, r, value, 307)
		} else {
			http.Error(w, "Error", http.StatusBadRequest)
		}

	} else {
		http.Error(w, "Error", http.StatusBadRequest)
	}
}
