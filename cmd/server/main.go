package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"study-event-api/internal/server"
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))
}

func main() {
	srv := server.NewServer()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
