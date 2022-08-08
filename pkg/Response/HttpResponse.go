package response

import (
	"Mach/pkg/logger"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Test us the result of Testcases
type TestCaseResult struct {
	Description string //Whats gone wrong
	Result      bool   // Pass or fail
}

func newTestCaseResult(Description string, Result bool) *TestCaseResult {

	tcr := TestCaseResult{Description: Description}
	tcr.Result = Result
	return &tcr
}

func ResponceMain(testcase interface{}, responce *http.Response) *TestCaseResult {

	Responces := testcase.(map[string]interface{})["Responces"]
	StatusCode := Responces.(map[string]interface{})["StatusCode"].(int)
	result := checkStatusCode(StatusCode, responce.StatusCode)
	if !result {
		return newTestCaseResult("Status Code Not Matched", false)
	}
	Body := Responces.(map[string]interface{})["Body"]
	respObj := ResponcetoObject(responce)
	return bodyContains(Body.(map[string]interface{})["Contains"], respObj)

}

func checkStatusCode(statusCode int, responceStatusCode int) bool {
	if statusCode == responceStatusCode {
		return true
	} else {
		return false
	}

}

func bodyContains(Contains interface{}, respObj interface{}) *TestCaseResult {

	if Contains.(map[string]interface{})["Type"] == "List" {

		switch respObj.(type) {
		case []interface{}:
			return list(Contains, respObj)
		default:
			return newTestCaseResult("No List Found", false)
		}

	} else if Contains.(map[string]interface{})["Type"] == "Object" {
		switch respObj.(type) {
		case map[string]interface{}:
			return object(Contains, respObj)
		default:
			return newTestCaseResult("No Obj Found", false)
		}

	} else {
		return newTestCaseResult("Passed", true)
	}

}

func list(Contains interface{}, respObj interface{}) *TestCaseResult {
	Lenght := Contains.(map[string]interface{})["Lenght"]
	if Lenght.(map[string]interface{})["Equal"] != nil {
		leng := Lenght.(map[string]interface{})["Equal"]
		if !Equal(leng.(int), len(respObj.([]interface{}))) {
			return newTestCaseResult("List Lenght not equal", false)
		}
	}
	InType := Contains.(map[string]interface{})["InType"]
	if InType != nil {
		return listValue(InType, respObj)
	}

	return newTestCaseResult("Passed", true)
}

func listValue(InType interface{}, respObj interface{}) *TestCaseResult {
	intype := InType.([]interface{})[0]

	for key, value := range intype.(map[interface{}]interface{}) {
		if len(intype.(map[interface{}]interface{})) < key.(int) {
			return newTestCaseResult("List item not found", false)
		}
		obj := respObj.([]interface{})[key.(int)]
		if obj == nil {
			return newTestCaseResult("List item not found", false)
		}
		switch value.(type) {
		case map[string]interface{}:

			res := bodyContains(value.(map[string]interface{})["Contains"], obj)
			if !res.Result {
				return res
			}

		default:

			if !ListChecks(value, obj) {
				return newTestCaseResult("List item not matched", false)
			}
		}

	}
	return newTestCaseResult("Passed", true)
}

func object(Contains interface{}, respObj interface{}) *TestCaseResult {
	Lenght := Contains.(map[string]interface{})["Lenght"]
	if Lenght.(map[string]interface{})["Equal"] != nil {
		leng := Lenght.(map[string]interface{})["Equal"]
		if !Equal(leng.(int), len(respObj.(map[string]interface{}))) {
			return newTestCaseResult("Object Lenght not equal", false)
		}
	}
	InType := Contains.(map[string]interface{})["InType"]

	if InType != nil {
		return objValue(InType, respObj)
	}
	return newTestCaseResult("Passed", true)
}

func objValue(InType interface{}, respObj interface{}) *TestCaseResult {
	intype := InType.([]interface{})[0]

	for key, value := range intype.(map[string]interface{}) {
		obj := respObj.(map[string]interface{})[key]
		if obj == nil {
			return newTestCaseResult("Object item not found", false)
		}

		val := value.(map[string]interface{})["Contains"]
		if val != nil {
			res := bodyContains(value, obj)
			if !res.Result {
				return res
			}
		} else {
			if !objChecks(value, obj) {
				return newTestCaseResult("List item not matched", false)
			}
		}

	}
	return newTestCaseResult("Passed", true)
}

func Equal[T int | string](yml T, resp T) bool {

	return yml == resp
}

func ResponcetoObject(resp *http.Response) interface{} {

	defer resp.Body.Close()
	resbody, _ := ioutil.ReadAll(resp.Body)

	var responceObj interface{}
	err := json.Unmarshal(resbody, &responceObj)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	return responceObj
}

func objChecks(ymlval interface{}, resobj interface{}) bool {

	val := ymlval.(map[string]interface{})["Equal"]
	if val != nil {
		return Equal(val.(string), resobj.(string))
	}
	return true
}

func ListChecks(ymlval interface{}, resobj interface{}) bool {

	return Equal(ymlval.(string), resobj.(string))

}
