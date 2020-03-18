package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/store", storeHandler)
	http.HandleFunc("/storage", storageHandler)

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
