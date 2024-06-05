package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hey", func(w http.ResponseWriter, r *http.Request) {
		handleHeyRequest(w, r)
	})

	port := getEnvOrDefault("PORT", "8080")
	fmt.Printf("Service 2 listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleHeyRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Service 2! Path: %s", r.URL.Path)
}

func getEnvOrDefault(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}
