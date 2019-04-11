package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name, _ := os.Hostname()
	fmt.Fprintf(w, "Hello there, enemies. I am %s", name)
}

func main() {
	fmt.Println("Starting!")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
