package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Starter applikasjon")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
