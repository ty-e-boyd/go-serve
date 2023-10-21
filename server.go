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

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		_, err2 := fmt.Fprintf(w, "Unable to parse form. %v", err)
		if err2 != nil {
			return
		}
		return
	}

	_, err := fmt.Fprintf(w, "POST was successful")
	if err != nil {
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	_, err = fmt.Fprintf(w, "Name: %s", name)
	if err != nil {
		return
	}

	_, err = fmt.Fprintf(w, "Address: %s", address)
	if err != nil {
		return
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
