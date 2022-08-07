package main

import (
	"fmt"
	"io/ioutil"
	"reflect"

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

		ResponceMain(testcase)
		//pkg.Request(testcase, RequestURL, HTTPmethods)
		//Responces := testcase.(map[string]interface{})["Responces"]
		// fmt.Println(Responces)

	}

	//fmt.Print(RequestURL)

}

func ResponceMain(testcase interface{}) {

	Responces := testcase.(map[string]interface{})["Responces"]
	StatusCode := Responces.(map[string]interface{})["StatusCode"]
	fmt.Println(StatusCode)
	Body := Responces.(map[string]interface{})["Body"]
	bodyContains(Body.(map[string]interface{})["Contains"])

}

func bodyContains(Contains interface{}) {

	if Contains.(map[string]interface{})["Type"] == "List" {
		list(Contains)
	} else if Contains.(map[string]interface{})["Type"] == "Object" {
		object()
	} else {
		fmt.Println("error: no body")
	}

}

func list(Contains interface{}) {
	Lenght := Contains.(map[string]interface{})["Lenght"]
	if Lenght.(map[string]interface{})["Equal"] != nil {
		Equal(Lenght)
	}
	InType := Contains.(map[string]interface{})["InType"]
	if InType != nil {
		listObj(InType)
	}

}

func listObj(InType interface{}) {
	intype := InType.([]interface{})[0]
	//fmt.Println(intype)

	for key, value := range intype.(map[interface{}]interface{}) {

		switch value.(type) {
		case map[string]interface{}:
			fmt.Println(key, reflect.TypeOf(value))
			bodyContains(value.(map[string]interface{})["Contains"])
			//fmt.Println("Integer:", value.(map[string]interface{})["Contains"])
		default:
			fmt.Println(key, reflect.TypeOf(value))
		}

	}
}

func object() {
	fmt.Println("obj")
}

func Equal(Lenght interface{}) {
	fmt.Println(Lenght)
}
