package covidstats

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
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

	yesterdayURL := "https://www.who.int/docs/default-source/coronaviruse/situation-reports/20200329-sitrep-69-covid-19.pdf?sfvrsn=8d6620fa_4"
	todayURL := "https://www.who.int/docs/default-source/coronaviruse/situation-reports/20200330-sitrep-70-covid-19.pdf?sfvrsn=7e0fe3f8_4"
	var statistics []DayStatistics

	if err := downloadFile("yesterday.pdf", yesterdayURL); err != nil {
		panic(err)
	}

	if err := downloadFile("today.pdf", todayURL); err != nil {
		panic(err)
	}

	todayRead, err := readPlainTextFromPDF("today.pdf")
	if err != nil {
		panic(err)
	}

	yesterdayRead, err := readPlainTextFromPDF("yesterday.pdf")
	if err != nil {
		panic(err)
	}

	yesterdaySplitted := strings.Fields(yesterdayRead)
	todaySplitted := strings.Fields(todayRead)
	yesterdayPosition := 0

	for i := 0; i < len(yesterdaySplitted); i++ {
		if yesterdaySplitted[i] == "Norway" {
			yesterdayPosition = i
			break
		}
	}

	//	fmt.Println(todaySplitted)
	for i := 0; i < len(todaySplitted); i++ {
		if todaySplitted[i] == "Norway" {

			statsToday := DayStatistics{InfectedTotal: todaySplitted[i+1], InfectedNew: todaySplitted[i+2], DeathsTotal: todaySplitted[i+3], DeathsNew: todaySplitted[i+4]}
			statsYesterday := DayStatistics{InfectedTotal: yesterdaySplitted[yesterdayPosition+1], InfectedNew: yesterdaySplitted[yesterdayPosition+2], DeathsTotal: yesterdaySplitted[yesterdayPosition+3], DeathsNew: yesterdaySplitted[yesterdayPosition+4]}

			statistics = append(statistics, statsToday)
			statistics = append(statistics, statsYesterday)

			fmt.Println("Norge har: ")
			break
		}
	}

	return statistics
}

func searchURL(url string) DayStatistics {

	if err := downloadFile("yesterday.pdf", yesterdayURL); err != nil {
		panic(err)
	}

	return stats

}

// Hentet fra: https://siongui.github.io/2018/09/21/go-read-plain-text-in-pdf-file/
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

// DownloadFile laster ned pdffilen
// Hentet fra: https://golangcode.com/download-a-file-from-a-url/
func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
