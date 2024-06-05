package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	isService1 := os.Getenv("SERVICE_TYPE") == "service1"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/srv2") {
			// Handle requests for /srv2/*
			handleSrv2Request(w, r)
		} else {
			// Handle requests for /
			if isService1 {
				// Redirect or forward requests from / to /srv2
				http.Redirect(w, r, "/srv2"+r.URL.Path, http.StatusTemporaryRedirect)
			} else {
				handleRootRequest(w, r)
			}
		}
	})

	port := getEnvOrDefault("PORT", "8080")
	fmt.Printf("Service listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleRootRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from /! Path: %s", r.URL.Path)
}

func handleSrv2Request(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from /srv2! Path: %s", r.URL.Path)
}

func getEnvOrDefault(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}