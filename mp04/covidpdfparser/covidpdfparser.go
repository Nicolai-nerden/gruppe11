package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/ledongthuc/pdf"
)

func main() {

	yesterday := "https://www.who.int/docs/default-source/coronaviruse/situation-reports/20200327-sitrep-67-covid-19.pdf?sfvrsn=b65f68eb_4"
	today := "https://www.who.int/docs/default-source/coronaviruse/situation-reports/20200329-sitrep-69-covid-19.pdf?sfvrsn=8d6620fa_4"

	if err := downloadFile("yesterday.pdf", yesterday); err != nil {
		panic(err)
	}

	if err := downloadFile("today.pdf", today); err != nil {
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

			fmt.Println("Norge har: ")
			fmt.Println("Smittede totalt:", todaySplitted[i+1], "(", yesterdaySplitted[yesterdayPosition+1], "i går)")
			fmt.Println("Smittede siste døgn:", todaySplitted[i+2], "(", yesterdaySplitted[yesterdayPosition+2], "i går)")
			fmt.Println("Dødsfall totalt:", todaySplitted[i+3], "(", yesterdaySplitted[yesterdayPosition+3], "i går)")
			fmt.Println("Dødsfall siste døgn:", todaySplitted[i+4], "(", yesterdaySplitted[yesterdayPosition+4], "i går)")
			break
		}
	}

	fmt.Println("ferdig")
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
