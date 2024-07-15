package main

import (
	"fmt"
	"net/http"
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
	key := r.PathValue("shortURL")

	if key == "" {
		app.notFound(w)
		return
	}

	originalURL, err := app.URLmodel.Get(key)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
