package report

import (
	"bufio"
	"bytes"
	"html/template"
	"os"
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

	tmpl := template.Must(template.ParseFiles("templates/report.gohtml"))

	var processed bytes.Buffer
	tmpl.Execute(&processed, report)

	//currentTime := time.Now().Format("yyyy-MM-dd'T'HH:mm:ss")

	outputPath := "reports/Report.html "
	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(processed.String())
	w.Flush()
}
