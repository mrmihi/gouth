package main

import (
	"fmt"
	"log"
	"net/http"
)

func respond() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "All GOOD!")
		if err != nil {
			return
		}
	})
}

func handleRequests() {
	http.Handle("/", respond())
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("server")
	handleRequests()
}
