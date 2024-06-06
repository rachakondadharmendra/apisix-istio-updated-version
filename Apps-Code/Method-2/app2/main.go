package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/srv2/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hey I'm service 2 ....")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Service 2 listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}