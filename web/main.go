package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, world!<h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>About me<h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("starting server...")
	http.ListenAndServe(":8080", nil)
}
