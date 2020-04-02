package covidanalytics

import (
	"bytes"
	"fmt"
	"gruppe11/mp04/latestreportlinks"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ledongthuc/pdf"
)

// DayStatistics inneholder informasjonen til ulike dager.
type DayStatistics struct {
	InfectedTotal string
	InfectedNew   string
	DeathsTotal   string
	DeathsNew     string
}

var fileNum int

// GetStatistics henter ut statistikkene fra i går og idag.
func GetStatistics() []DayStatistics {

	fmt.Println("Vent mens statistikken analyseres... \nDet vil snart åpne seg en fane i nettleseren din med satistikken.")
	todayURL, yesterdayURL := latestreportlinks.FindReportLinks("https://www.who.int/emergencies/diseases/novel-coronavirus-2019/situation-reports")

	today := searchURL(todayURL)
	yesterday := searchURL(yesterdayURL)

	statistics := []DayStatistics{today, yesterday}

	return statistics
}

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

	fileNum++

	stats := DayStatistics{InfectedTotal: pdfClean[norwayPos+1], InfectedNew: pdfClean[norwayPos+2], DeathsTotal: pdfClean[norwayPos+3], DeathsNew: pdfClean[norwayPos+4]}

	return stats
}

// readPlainTextFromPDF er Ikke selvlaget. Hentet fra: https://siongui.github.io/2018/09/21/go-read-plain-text-in-pdf-file/
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

	resp, err := http.Get(url)
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
