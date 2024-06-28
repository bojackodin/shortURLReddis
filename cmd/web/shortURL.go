package main

import (
	"fmt"
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

	shortKey := generateShortURL(originalURL)
	err := app.URLmodel.Add(originalURL, shortKey)
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
	originalURL, err := app.URLmodel.Get(key)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
