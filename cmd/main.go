package main

import (
	"Mach/pkg"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func main() {
	// test.read yml

	buf := readFile()
	//Map the read file[interface {}]interface {}Map to
	t := mapYmltoInterface(buf)

	performMach(t)

}

func readFile() []byte {
	buf, err := ioutil.ReadFile("Apitest.yml")
	if err != nil {
		fmt.Print("error: Failed to read the file\n")
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

func performMach(t map[interface{}]interface{}) {

	fmt.Print(t["APITESTING"].(map[string]interface{})["Name"]) // Name
	fmt.Print("\n")
	RequestURL := t["APITESTING"].(map[string]interface{})["RequestURL"]
	fmt.Println(RequestURL)
	HTTPmethods := t["APITESTING"].(map[string]interface{})["HTTP-method"].(string)
	fmt.Println(HTTPmethods)
	testcases := t["APITESTING"].(map[string]interface{})["TestCases"].([]interface{})
	testcasesLen := len(testcases)
	for i := 0; i < testcasesLen; i++ {
		testcase := testcases[i]
		Name := testcase.(map[string]interface{})["Name"]
		fmt.Println(Name)
		pkg.Request(testcase, RequestURL, HTTPmethods)
		//Responces := testcase.(map[string]interface{})["Responces"]
		// fmt.Println(Responces)

	}

	//fmt.Print(RequestURL)

}
