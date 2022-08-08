package main

import (
	"Mach/pkg/ReadFiles"
	"Mach/pkg/Request"
	"Mach/pkg/Response"
	"fmt"
)

func main() {
	// test.read yml

	t := ReadFiles.ReadYml()

	performMach(t)

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

		responce := Request.Request(testcase, RequestURL, HTTPmethods)

		//fmt.Print(RequestURL)
		fmt.Println(responce.StatusCode)
		// vars, err := respObj.(map[string]interface{})["city"]
		// if err == true {
		// 	fmt.Print("error")
		// }

		result := Response.ResponceMain(testcase, responce)

		fmt.Println(result)
		//
		//Responces := testcase.(map[string]interface{})["Responces"]
		// fmt.Println(Responces)

	}

	//fmt.Print(RequestURL)

}
