package main

import (
	"fmt"
	"html"
	"log/slog"
	"net/http"
	"os"

	"github.com/jseiser/jls2/middleware"
)

func live() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("I am Alive"))
	})
}

func ready() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("I am Ready"))
	})
}

func not_found() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error: Handler for %s not found", html.EscapeString(r.URL.Path))
	})
}
func main() {
	slogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(slogger)

	mux := http.NewServeMux()

	mux.Handle("GET /live", middleware.HttpLogger(live()))

	mux.Handle("GET /ready", middleware.HttpLogger(ready()))

	mux.Handle("GET /", middleware.HttpLogger(not_found()))

	slogger.Info("Starting Service")
	http.ListenAndServe(":3333", mux)
}
