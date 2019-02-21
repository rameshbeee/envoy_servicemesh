package main

import (
	"fmt"
	"log"

	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello from service B\n")
}

func main() {
	http.HandleFunc("/service_b", handler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
