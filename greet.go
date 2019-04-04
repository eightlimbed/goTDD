package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet writes a greeting to a Writer
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

// MyGreeterHandler writes a greeting to a web client
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
