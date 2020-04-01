package main

import (
	"fmt"
	"gruppe11/mp04/covidstats"
	"net/http"

	"github.com/pkg/browser"
)

var opened = false

func indexHandler(w http.ResponseWriter, r *http.Request) {
	statistics := covidstats.GetStatistics()
	fmt.Fprintf(w, "<h1>")
	fmt.Fprintf(w, statistics[0].InfectedTotal)
	fmt.Fprintf(w, "</h1>")

}

func main() {
	if !opened {
		browser.OpenURL("http://127.0.0.1:8000/")
		opened = true
	}
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
