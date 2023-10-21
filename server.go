package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 - Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not available", http.StatusNotFound)
		return
	}

	_, err := fmt.Fprintf(w, "Hey boi.")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
