package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.Query().Get("val")))
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
