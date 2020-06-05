package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	renderer, err := NewRenderer()
	if err != nil {
		log.Fatal(fmt.Errorf("could not initialize renderer: %w", err))
		return
	}

	handler := newHandler(renderer)
	http.HandleFunc("/store", handler.store)
	http.HandleFunc("/storage", handler.storage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf(`
	Server started on localhost:%s
		/store   : store a new link
		/storage : get link stored in memory
	`, port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
