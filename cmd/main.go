package main

import (
	"Mach/pkg"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"gopkg.in/yaml.v3"
)

// Test us the result of Testcases
type TestCaseTests struct {
	Description string //Whats gone wrong
	Result      bool   // Pass or fail
}

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

		responce := pkg.Request(testcase, RequestURL, HTTPmethods)

		//fmt.Print(RequestURL)
		fmt.Println(responce.StatusCode)
		// vars, err := respObj.(map[string]interface{})["city"]
		// if err == true {
		// 	fmt.Print("error")
		// }

		result := ResponceMain(testcase, responce)

		fmt.Println(result)
		//
		//Responces := testcase.(map[string]interface{})["Responces"]
		// fmt.Println(Responces)

	}

	//fmt.Print(RequestURL)

}

func ResponceMain(testcase interface{}, responce *http.Response) bool {

	result := true
	Responces := testcase.(map[string]interface{})["Responces"]
	StatusCode := Responces.(map[string]interface{})["StatusCode"].(int)
	result = checkStatusCode(StatusCode, responce.StatusCode)
	if result == false {
		return false
	}
	Body := Responces.(map[string]interface{})["Body"]
	respObj := ResponcetoObject(responce)
	fmt.Println(reflect.TypeOf(respObj.([]interface{})[0]))
	return bodyContains(Body.(map[string]interface{})["Contains"], respObj)

}

func checkStatusCode(statusCode int, responceStatusCode int) bool {
	if statusCode == responceStatusCode {
		return true
	} else {
		return false
	}

}

func bodyContains(Contains interface{}, respObj interface{}) bool {

	if Contains.(map[string]interface{})["Type"] == "List" {

		switch respObj.(type) {
		case []interface{}:
			return list(Contains, respObj)
		default:
			return false
		}

	} else if Contains.(map[string]interface{})["Type"] == "Object" {
		switch respObj.(type) {
		case map[string]interface{}:
			return object(Contains, respObj)
		default:
			return false
		}

	} else {
		return false
	}

}

func list(Contains interface{}, respObj interface{}) bool {
	Lenght := Contains.(map[string]interface{})["Lenght"]
	if Lenght.(map[string]interface{})["Equal"] != nil {
		leng := Lenght.(map[string]interface{})["Equal"]
		if Equal(leng.(int), len(respObj.([]interface{}))) == false {
			return false
		}
	}
	InType := Contains.(map[string]interface{})["InType"]
	if InType != nil {
		return listValue(InType, respObj)
	}

	return true
}

func listValue(InType interface{}, respObj interface{}) bool {
	intype := InType.([]interface{})[0]
	//fmt.Println(intype)

	for key, value := range intype.(map[interface{}]interface{}) {
		obj := respObj.([]interface{})[key.(int)]
		if obj == nil {
			return false
		}
		switch value.(type) {
		case map[string]interface{}:
			fmt.Println(key, reflect.TypeOf(value))
			if bodyContains(value.(map[string]interface{})["Contains"], obj) == false {
				return false
			}
			//fmt.Println("Integer:", value.(map[string]interface{})["Contains"])
		default:
			fmt.Println(key, reflect.TypeOf(value))
			if ListChecks(value, obj) == false {
				return false
			}
		}

	}
	return true
}

func object(Contains interface{}, respObj interface{}) bool {
	Lenght := Contains.(map[string]interface{})["Lenght"]
	if Lenght.(map[string]interface{})["Equal"] != nil {
		leng := Lenght.(map[string]interface{})["Equal"]
		if Equal(leng.(int), len(respObj.(map[string]interface{}))) == false {
			return false
		}
	}
	InType := Contains.(map[string]interface{})["InType"]

	//fmt.Println(InType)
	if InType != nil {
		return objValue(InType, respObj)
	}
	return true
}

func objValue(InType interface{}, respObj interface{}) bool {
	intype := InType.([]interface{})[0]
	//fmt.Println(intype)

	for key, value := range intype.(map[string]interface{}) {
		obj := respObj.(map[string]interface{})[key]
		if obj == nil {
			return false
		}
		fmt.Println(key, reflect.TypeOf(value), obj)

		val := value.(map[string]interface{})["Contains"]
		if val != nil {
			if bodyContains(value, obj) == false {
				return false
			}
		} else {
			if objChecks(value, obj) == false {
				return false
			}
		}

	}
	return true
}

func Equal[T int | string](yml T, resp T) bool {
	//fmt.Println("equal", yml, resp)
	if yml == resp {
		return true
	}
	return false
}

func ResponcetoObject(resp *http.Response) interface{} {

	// fmt.Println(res.Status)
	defer resp.Body.Close()
	resbody, _ := ioutil.ReadAll(resp.Body)

	var responceObj interface{}
	err := json.Unmarshal(resbody, &responceObj)
	if err != nil {
		fmt.Println(err)
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
