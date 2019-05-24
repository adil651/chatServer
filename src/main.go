package main

import (
	"log"
	"net/http"
)

func main() {

	// Create a simple file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	log.Println("http server started on :8000")
	// Start the HTTP server
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
