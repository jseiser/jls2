package main

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	trace := string(debug.Stack())

	uri := r.RequestURI
	method := r.Method
	protocol := r.Proto
	remote_addr := r.RemoteAddr

	app.logger.Error(
		"Internal Service Error",
		slog.String("method", method),
		slog.String("protocol", protocol),
		slog.String("uri", uri),
		slog.String("remote_addr", remote_addr),
		slog.Any("error", err),
		slog.Any("trace", trace),
	)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
