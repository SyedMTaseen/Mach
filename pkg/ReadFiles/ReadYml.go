package readFiles

import (
	"Mach/pkg/logger"
	"Mach/pkg/macher"
	"Mach/pkg/report"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

func ReadYml(testCasesPath string) {

	var reports report.Report
	currentTime := time.Now()
	reports.ReportDate = currentTime.Format("01-02-2006 15:04:05 Monday")
	for _, s := range find(testCasesPath, ".yml") {
		buf := readFile(s)
		interf := mapYmltoInterface(buf)
		tests := macher.PerformMach(interf)
		reports.Tests = append(reports.Tests, tests)
	}

	report.CreateReport(reports)

}

func readFile(path string) []byte {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		logger.ErrorLogger.Println("error: Failed to read the file\n" + ":" + err.Error())
		return nil
	}
	return buf
}

func mapYmltoInterface(buf []byte) map[interface{}]interface{} {
	t := make(map[interface{}]interface{})
	err := yaml.Unmarshal(buf, &t)
	if err != nil {
		panic(err)
	}
	return t
}

func find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}
