package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world from golang")
		log.Printf("request to /")
	})

	http.ListenAndServe(":80", nil)
}
