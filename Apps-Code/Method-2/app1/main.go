// main.go for service1
package main

import (
	"fmt"
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
		// Redirect to service2 with /srv2 prefix
		http.Redirect(w, r, service2Addr+"/srv2"+r.URL.Path, http.StatusTemporaryRedirect)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Service 1 listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
