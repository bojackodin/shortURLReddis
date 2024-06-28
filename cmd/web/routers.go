package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/shorten/", app.createShortURL)
	mux.HandleFunc("/short/{shortURL}", app.getShortURL)

	return app.logRequest(mux)
}
