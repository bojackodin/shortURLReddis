package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HI")
}

func (app *application) createShortURL(w http.ResponseWriter, r *http.Request) {
	originalURL := r.URL.Query().Get("url")
	if originalURL == "" {
		app.notFound(w)
		return
	}

	shortKey := generateShortURL()
	err := app.URLmodel.Add(app.ctx, originalURL, shortKey)
	if err != nil {
		app.serverError(w, err)
		return
	}
	shortURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

	fmt.Fprintln(w, shortURL)
}

func (app *application) getShortURL(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	parts := strings.Split(path, "/")
	if len(parts) > 3 {
		app.notFound(w)
		return
	}

	key := parts[len(parts)-1]
	originalURL, err := app.URLmodel.Get(app.ctx, key)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}

func generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}
