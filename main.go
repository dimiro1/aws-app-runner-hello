package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "<h1>Hello AWS App Runner</h1>")
	})

	log.Printf("Starting the server on :%s...", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
