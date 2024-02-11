package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Minikube!\n")
	})

	// compose a server
	addr := ":8081"

	server := &http.Server{
		Addr: addr,
	}

	// run the server
	fmt.Printf("Running on port %s\n", addr)
	server.ListenAndServe()
}

