package macher

import (
	"Mach/pkg/report"
	"Mach/pkg/request"
	"Mach/pkg/response"
	"io/ioutil"
	"strconv"
)

func PerformMach(t map[interface{}]interface{}) report.Test {

	var tests report.Test
	tests.Discription = t["Mach"].(map[string]interface{})["Name"].(string)
	//fmt.Println("API Discription: " + t["Mach"].(map[string]interface{})["Name"].(string)) // Name
	RequestURL := t["Mach"].(map[string]interface{})["RequestURL"]
	//fmt.Println("Request URL " + RequestURL.(string))
	tests.Url = RequestURL.(string)
	HTTPmethods := t["Mach"].(map[string]interface{})["HTTP-method"].(string)
	//fmt.Println("Method: " + HTTPmethods)
	tests.Request = HTTPmethods
	testcases := t["Mach"].(map[string]interface{})["TestCases"].([]interface{})
	testcasesLen := len(testcases)
	for i := 0; i < testcasesLen; i++ {
		var results report.Result
		testcase := testcases[i]
		Name := testcase.(map[string]interface{})["Name"]
		results.Testcase = Name.(string)
		//fmt.Print("Test case : " + Name.(string) + " ")

		responce := request.Request(testcase, RequestURL, HTTPmethods)

		//fmt.Println(responce.StatusCode)

		result := response.ResponceMain(testcase, responce)

		if result.Result {
			results.Status = true
			//fmt.Println("PASSED")
		} else {
			results.Status = false
			var reasons report.Reason
			reasons.Description = result.Description
			// fmt.Println("FAILED")
			// fmt.Println("Reason: " + result.Description)
			resbody, _ := ioutil.ReadAll(responce.Body)
			reasons.Code = strconv.Itoa(responce.StatusCode)
			//fmt.Println("Status: " + strconv.Itoa(responce.StatusCode))
			reasons.Response = string(resbody)
			results.Reasons = reasons
			//fmt.Println("Response: " + string(resbody))
		}

		tests.Results = append(tests.Results, results)
		//fmt.Println(result)

	}
	return tests

}
