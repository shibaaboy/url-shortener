package handler

import (
	"io"
	"net/http"
)

var storage = make(map[string]string)
var https = "https://"
var id = "EwHXdJfB"
var errorMessage = "Sorry, error..."

func MainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		body, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, errorMessage, http.StatusInternalServerError) //500
			return
		}

		originalURL := string(body)
		storage[id] = originalURL

		w.WriteHeader(201)
		w.Write([]byte(https + r.Host + "/" + id))

	case http.MethodGet:
		currId := r.URL.Path
		currIdWithoutSlash := currId[1:]

		value, ok := storage[currIdWithoutSlash]

		if ok {
			http.Redirect(w, r, value, http.StatusTemporaryRedirect) //307
		} else {
			http.Error(w, errorMessage, http.StatusBadRequest) //400
		}

	default:
		http.Error(w, errorMessage, http.StatusMethodNotAllowed) //405
	}
}
