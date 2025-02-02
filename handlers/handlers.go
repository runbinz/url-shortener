package handlers

import (
	"fmt"
	"net/http"
)

var urlMap = make(map[string]string)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("url")
	if longURL == "" {
		http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
		return
	}

	shortCode := fmt.Sprintf("%d", len(urlMap)+1)
	urlMap[shortCode] = longURL
	w.Write([]byte(fmt.Sprintf("Short URL: http://localhost:8080/%s", shortCode)))
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:]
	longURL, exists := urlMap[shortCode]
	if !exists {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}
