package latestreportlinks

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// FindReportLinks søker gjennom WHO sin rapport oversikt og finner linken til de filene som er ønsket.
func FindReportLinks(url string) (string, string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	latestReportNum := GetLatestNum(string(bytes))
	choppedHTML := strings.Fields(string(bytes))

	today := getURLOf(choppedHTML, latestReportNum)
	yesterday := getURLOf(choppedHTML, latestReportNum-1)

	return today, yesterday

}

// GetLatestNum finner siste rapportnr som er lagt ut.
func GetLatestNum(html string) int {
	latestReportNum := 60 // vi vet at det er minst 60 rapporter ute

	for strings.Contains(html, "Situation report - "+strconv.Itoa(latestReportNum)) {
		latestReportNum++
	}
	latestReportNum += 2

	return latestReportNum
}

// getURLOf finner rapporten med nummeret digg i parameteret og filterer ut linken.
func getURLOf(choppedHTML []string, latestReportNum int) string {

	sitRepNr := 20

	for i := 0; i < len(choppedHTML); i++ {
		if choppedHTML[i] == "report" && choppedHTML[i+1] == "-" && (choppedHTML[i+2] == strconv.Itoa(latestReportNum)+"</strong></a><br" || choppedHTML[i+2] == strconv.Itoa(latestReportNum)+"</a></strong><br") {
			fmt.Println("Sammenligner rapport", latestReportNum)
			sitRepNr = i
			break
		}
	}

	urlTag := choppedHTML[sitRepNr-1]
	start := 0
	end := 0

	for i := 0; i < len(urlTag); i++ {
		if start == 0 && []byte(urlTag[i : i+1])[0] == 34 {
			start = i + 1
		} else if []byte(urlTag[i : i+1])[0] == 34 {
			end = i
			break
		}
	}

	return "https://www.who.int" + urlTag[start:end]
}
