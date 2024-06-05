package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	isService1 := os.Getenv("SERVICE_TYPE") == "service1"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if isService1 {
			// Redirect all requests from / to /hey in app2
			http.Redirect(w, r, "/hey"+r.URL.Path, http.StatusTemporaryRedirect)
		} else {
			handleRootRequest(w, r)
		}
	})

	port := getEnvOrDefault("PORT", "8080")
	fmt.Printf("Service 1 listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleRootRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Service 1! Path: %s", r.URL.Path)
}

func getEnvOrDefault(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}
