package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	models "github.com/jseiser/jls2/internal"
)

type application struct {
	logger *slog.Logger
	todos  *models.Todos
}

func main() {
	addr := flag.String("addr", ":3333", "HTTP network address")

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelInfo,
	}))

	app := &application{
		logger: logger,
		todos:  &models.Todos{},
	}

	logger.Info("starting server", slog.String("addr", ":4000"))
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
