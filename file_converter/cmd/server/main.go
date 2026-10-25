package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Shashwat8127/Project/file_converter/api"
	"github.com/Shashwat8127/Project/file_converter/handlers"
)

func main() {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           api.Handler(&handlers.Handler{}),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("file converter listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
