package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the service name from the request's Host header
		serviceName := r.Host

		// Construct the URL for the /hey endpoint of the service within the cluster domain
		redirectURL := fmt.Sprintf("http://%s/hey", serviceName)
		fmt.Printf("Received URL: %s\n", r.URL.String())
		fmt.Printf("Redirect URL: %s\n", redirectURL)

		// Redirect the request to the constructed URL
		http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
	})

	port := "8080"
	fmt.Printf("Service 1 listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
