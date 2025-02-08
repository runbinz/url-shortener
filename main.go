package main

import (
	"fmt"
	"net/http"
	"url-shortener/handlers"
)

func main() {
	// create instance of URLShortener
	shortener := &handlers.URLShortener{
		Urls: make(map[string]string),
	}

	// set up HTTP handlers
	http.HandleFunc("/", serveForm)
	http.HandleFunc("/shorten", shortener.HandleShorten)
	http.HandleFunc("/short/", shortener.HandleRedirect)

	fmt.Println("URL Shortener is running on :8080")
	http.ListenAndServe(":8080", nil)
}

// serveForm serves the HTML form for URL shortening
func serveForm(w http.ResponseWriter, r *http.Request) {
	html := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>URL Shortener</title>
        </head>
        <body>
            <h2>URL Shortener</h2>
            <form method="post" action="/shorten">
                <input type="text" name="url" placeholder="Enter a URL">
                <input type="submit" value="Shorten">
            </form>
        </body>
        </html>
    `
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}
