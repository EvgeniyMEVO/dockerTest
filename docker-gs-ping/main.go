package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/docker", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "<h1>Hello from Docker</h1>")
	})

	http.ListenAndServe("localhost:8080", nil)
}
