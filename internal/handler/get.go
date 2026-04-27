package handler

import "net/http"

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
