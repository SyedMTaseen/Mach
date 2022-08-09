package report

import (
	"bufio"
	"bytes"
	"html/template"
	"os"
	"time"
)

type Result struct {
	Testcase string
	Status   bool
	Reasons  Reason
}

type Reason struct {
	Description string
	Code        string
	Response    string
}

type Test struct {
	Discription string
	Url         string
	Request     string
	Results     []Result
}

type Report struct {
	ReportDate string
	Tests      []Test
}

func CreateReport(report Report) {

	tmpl := template.Must(template.ParseFiles("././templates/report.gohtml"))

	var processed bytes.Buffer
	tmpl.Execute(&processed, report)

	currentTime := time.Now()

	outputPath := "././reports/Report " + currentTime.Format("01-02-2006 15:04:05 Monday") + ".html"
	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(processed.String())
	w.Flush()
}
