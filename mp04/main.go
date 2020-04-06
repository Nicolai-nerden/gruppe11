package main

import (
	"fmt"
	"modul4/covidanalytics"
	"modul4/style"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/teamwork/reload"
)

var opened = false
var statistics = covidanalytics.GetStatistics()
var printQueue = []string{
	statistics[0].InfectedTotal, statistics[1].InfectedTotal,
	statistics[0].InfectedNew, statistics[1].InfectedNew,
	statistics[0].DeathsTotal, statistics[1].DeathsTotal,
	statistics[0].DeathsNew, statistics[1].DeathsNew,
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, style.Style)
	fmt.Fprintf(w, style.MainStart)
	fmt.Fprintf(w, "<tr>")

	lineShift := 0

	for i := 0; i < 4; i++ {
		today, _ := strconv.Atoi(printQueue[lineShift])
		yesterday, _ := strconv.Atoi(printQueue[lineShift+1])
		dif := today - yesterday
		difConv := math.Abs(float64(dif))

		fmt.Fprintf(w, "<tr>")
		for g := 0; g < 3; g++ {
			if g == 2 {
				if dif < 0 {
					fmt.Fprintf(w, "<td style=\"color: green;\">"+strconv.Itoa(int(difConv))+" ↓</td>")
				} else {
					fmt.Fprintf(w, "<td style=\"color: red;\">"+strconv.Itoa(int(difConv))+" ↑</td>")
				}
			} else {
				fmt.Fprintf(w, "<td>"+printQueue[g+lineShift]+"</td>")
			}

		}
		fmt.Fprintf(w, "</tr>")
		lineShift += 2
	}
	fmt.Fprintf(w, style.MainEnd)

}

func main() {

	// Fjern kommentarer her hvis du vil at nettleseren skal åpnes automatisk
	// if !opened {
	// 	browser.OpenURL("http://127.0.0.1:8081/")
	// 	opened = true
	// }

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))

	timer1 := time.NewTimer(1 * time.Hour)

	<-timer1.C
	reload.Exec()

}
