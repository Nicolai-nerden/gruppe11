package covidanalytics

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"modul4/latestreportlinks"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ledongthuc/pdf"
)

// DayStatistics inneholder statistikken for hver dag.
type DayStatistics struct {
	InfectedTotal string
	InfectedNew   string
	DeathsTotal   string
	DeathsNew     string
}

var fileNum int

// GetStatistics henter ut statistikkene fra i går og idag.
func GetStatistics() []DayStatistics {

	fmt.Println("Laster inn rapporter... Vennligst vent. \nDet vil snart åpne seg en fane i nettleseren din med satistikken.")
	todayURL, yesterdayURL := latestreportlinks.FindReportLinks("https://www.who.int/emergencies/diseases/novel-coronavirus-2019/situation-reports")

	today := searchURL(todayURL)
	yesterday := searchURL(yesterdayURL)

	statistics := []DayStatistics{today, yesterday}

	return statistics
}

// searchURL søker etter norges statistikk.
func searchURL(url string) DayStatistics {

	if err := downloadFile(strconv.Itoa(fileNum)+".pdf", url); err != nil {
		panic(err)
	}

	pdfConverted, err := readPlainTextFromPDF(strconv.Itoa(fileNum) + ".pdf")
	if err != nil {
		panic(err)
	}
	pdfClean := strings.Fields(pdfConverted) //renser filen og deler opp ordene i en slice.
	norwayPos := 0

	for i := 0; i < len(pdfClean); i++ {
		if pdfClean[i] == "Norway" {
			norwayPos = i
			break
		}
	}
	if fileNum >= 1 {
		fileNum = 0
	} else {
		fileNum++
	}

	stats := DayStatistics{InfectedTotal: pdfClean[norwayPos+1], InfectedNew: pdfClean[norwayPos+2], DeathsTotal: pdfClean[norwayPos+3], DeathsNew: pdfClean[norwayPos+4]}

	return stats
}

// readPlainTextFromPDF parser pdf til tekst i form av en string.
func readPlainTextFromPDF(pdfpath string) (text string, err error) {
	f, r, err := pdf.Open(pdfpath)
	defer f.Close()
	if err != nil {
		return
	}

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return
	}

	buf.ReadFrom(b)
	text = buf.String()
	return
}

// DownloadFile laster ned en fil via url.
func downloadFile(filepath string, url string) error {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
