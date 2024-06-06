package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		service2Addr := os.Getenv("SERVICE2_ADDR")
		if service2Addr == "" {
			service2Addr = "http://localhost:8080" // Default service2 address if not set
		}

		// Construct the URL for the service2 endpoint
		service2URL := service2Addr + "/srv2" + r.URL.Path

		// Create a new HTTP request to service2
		req, err := http.NewRequest("GET", service2URL, nil)
		if err != nil {
			http.Error(w, "Failed to create request to service2", http.StatusInternalServerError)
			return
		}

		// Perform the request to service2
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Failed to call service2", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Copy the response from service2 to the original response
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response from service2", http.StatusInternalServerError)
			return
		}

		// Copy the status code and headers from service2
		w.WriteHeader(resp.StatusCode)
		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		// Write the body from service2
		w.Write(body)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Service 1 listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
