package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Printf("Hello")
	})

	http.ListenAndServe(":9090", nil)
}
