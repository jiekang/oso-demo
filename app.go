package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name, _ := os.Hostname()

	// TODO: log request and print to stdout
	fmt.Fprintf(w, "Hello there enemies. I am %s", name)
}

func main() {
	fmt.Println("Starting!")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
