package main

import (
	"fmt"
	"net/http"
	"url-shortener/handlers"
)

func main() {
	http.HandleFunc("/shorten", handlers.ShortenURL)
	http.HandleFunc("/", handlers.RedirectURL)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
