package main

import (
	"log"
	"net/http"
	"study-event-api/internal/server"
)

func main() {
	srv := server.NewServer()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
