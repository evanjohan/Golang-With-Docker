package main

import (
	"fmt"
	"log"
	"net/http"
)

func docker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Images Golang Apps success in Docker version 1.0!")
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome BasiCoding!")
}

func main() {
	http.HandleFunc("/", docker)
	http.HandleFunc("/welcome", welcome)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
