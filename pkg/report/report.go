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

var tmpl *template.Template

func CreateReport() {

	res := Reason{Description: "hello", Code: "500", Response: "{\"ERROR\":\"server error\"}"}

	result1 := Result{Testcase: "basic test", Status: true}

	result2 := Result{Testcase: "basic test", Status: false, Reasons: res}

	tmpl = template.Must(template.ParseFiles("././templates/report.gohtml"))

	data := Report{
		ReportDate: "9 August 2022",
		Tests: []Test{
			{Discription: "Basic testing", Url: "https://www.figma.com/", Request: "GET", Results: []Result{result1, result2}},
			//	{Discription: "Basic testing", Url: "https://www.figma.com/", request: "GET", results: []Result{result1, result2}},
		},
	}

	var processed bytes.Buffer
	tmpl.Execute(&processed, data)

	currentTime := time.Now()

	outputPath := "././reports/Report " + currentTime.Format("01-02-2006 15:04:05 Monday") + ".html"
	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(processed.String())
	w.Flush()
}
