package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// manage mapping between original urls and shortened urls, urls is a map with key as the shortened url and value as the original url
type URLShortener struct {
	Urls map[string]string
}

// HandleShorten method for URLShortener struct -> handles post req, validates url input, generates unique short key
func (us *URLShortener) HandleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// get the original url from the form
	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
		return
	}

	// validate the original url
	_, err := url.ParseRequestURI(originalURL)
	if err != nil {
		http.Error(w, "Invalid 'url' parameter", http.StatusBadRequest)
		return
	}

	// generate a unique shortened key for the original url
	shortKey := generateShortkey()
	us.Urls[shortKey] = originalURL

	fmt.Printf("Generated short key: %s, Original URL: %s\n", shortKey, originalURL)

	// Construct short url
	shortURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

	// render html response w/ short url
	w.Header().Set("Content-Type", "text/html") // indicates html response to browser
	responseHTML := fmt.Sprintf(`
		<h2>URL Shortener</h2>
        <p>Original URL: %s</p>
        <p>Shortened URL: <a href="%s">%s</a></p>
        <form method="post" action="/shorten">
            <input type="text" name="url" placeholder="Enter a URL">
            <input type="submit" value="Shorten">
        </form>
	`, originalURL, shortURL, shortURL)
	fmt.Fprint(w, responseHTML)
}

// URL redirection
func (us *URLShortener) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortKey := r.URL.Path[len("/short/"):] // index from start of short key onwards
	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	// get the original url from the urls map using the short key,
	originalURL, found := us.Urls[shortKey] //found is a boolean indicating if the key is present in the map
	if !found {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	// redirect to the original url
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}

// generate short keys
func generateShortkey() string {
	// characters allowed
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6
	// random source and rng
	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)
	// create random byte slice to hold the short key
	shortKey := make([]byte, keyLength)
	// generate random key
	for i := range shortKey {
		// picks random from charset
		shortKey[i] = charset[randGenerator.Intn(len(charset))]
	}
	return string(shortKey)
}
