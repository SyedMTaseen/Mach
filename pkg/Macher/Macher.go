package macher

import (
	"Mach/pkg/request"
	"Mach/pkg/response"
	"fmt"
	"io/ioutil"
	"strconv"
)

func PerformMach(t map[interface{}]interface{}) {

	fmt.Println("API Discription: " + t["Mach"].(map[string]interface{})["Name"].(string)) // Name
	RequestURL := t["Mach"].(map[string]interface{})["RequestURL"]
	fmt.Println("Request URL " + RequestURL.(string))
	HTTPmethods := t["Mach"].(map[string]interface{})["HTTP-method"].(string)
	fmt.Println("Method: " + HTTPmethods)
	testcases := t["Mach"].(map[string]interface{})["TestCases"].([]interface{})
	testcasesLen := len(testcases)
	for i := 0; i < testcasesLen; i++ {
		testcase := testcases[i]
		Name := testcase.(map[string]interface{})["Name"]
		fmt.Print("Test case : " + Name.(string) + " ")

		responce := request.Request(testcase, RequestURL, HTTPmethods)

		//fmt.Println(responce.StatusCode)

		result := response.ResponceMain(testcase, responce)

		if result.Result {
			fmt.Println("PASSED")
		} else {
			fmt.Println("FAILED")
			fmt.Println("Reason: " + result.Description)
			resbody, _ := ioutil.ReadAll(responce.Body)
			fmt.Println("Status: " + strconv.Itoa(responce.StatusCode))
			fmt.Println("Response: " + string(resbody))
		}
		//fmt.Println(result)

	}

}
