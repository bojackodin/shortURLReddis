package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func generateShortURL(originalURL string) string {
	const keyLength = 7

	hash := sha256.Sum256([]byte(originalURL))
	shortURL := base64.URLEncoding.EncodeToString(hash[:])

	return shortURL[:keyLength]
}

func validateURL(u string) bool {
	re := regexp.MustCompile(`^(http|https)://`)
	if !re.MatchString(u) {
		return false
	}

	parsedURL, err := url.ParseRequestURI(u)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}

	return true
}
