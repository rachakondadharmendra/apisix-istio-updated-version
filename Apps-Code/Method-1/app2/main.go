package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hey", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hey I'm service 2 ....")
	})

	port := "8080"
	fmt.Printf("Service 2 listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
