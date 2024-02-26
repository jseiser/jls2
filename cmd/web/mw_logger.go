package main

import (
	"log/slog"
	"net/http"
	"time"
)

type responseData struct {
	status int
	size   int
}

type loggingResponseWriter struct {
	http.ResponseWriter
	responseData *responseData
}

func (lw *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := lw.ResponseWriter.Write(b)
	lw.responseData.size += size
	return size, err
}

func (lw *loggingResponseWriter) WriteHeader(statusCode int) {
	lw.ResponseWriter.WriteHeader(statusCode)
	lw.responseData.status = statusCode
}

func (app *application) httpLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}

		lw := loggingResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}

		uri := r.RequestURI
		method := r.Method
		protocol := r.Proto
		remote_addr := r.RemoteAddr

		next.ServeHTTP(&lw, r)
		status := lw.responseData.status
		size := lw.responseData.size
		duration := time.Since(start)

		app.logger.Info(
			"incoming request",
			slog.Int("status", status),
			slog.String("method", method),
			slog.String("protocol", protocol),
			slog.String("uri", uri),
			slog.String("remote_addr", remote_addr),
			slog.Duration("duration", duration),
			slog.Int("size", size),
		)

	})
}
