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

		ResponceMain(testcase, responce)
		//
		//Responces := testcase.(map[string]interface{})["Responces"]
		// fmt.Println(Responces)

	}

	//fmt.Print(RequestURL)

}

func ResponceMain(testcase interface{}, responce *http.Response) {

	Responces := testcase.(map[string]interface{})["Responces"]
	StatusCode := Responces.(map[string]interface{})["StatusCode"].(int)
	checkStatusCode(StatusCode, responce.StatusCode)
	Body := Responces.(map[string]interface{})["Body"]
	respObj := ResponcetoObject(responce)
	fmt.Println(reflect.TypeOf(respObj.([]interface{})[0]))
	bodyContains(Body.(map[string]interface{})["Contains"], respObj)

}

func checkStatusCode(statusCode int, responceStatusCode int) {
	if statusCode == responceStatusCode {
		fmt.Println("StaTusCode match")
	} else {
		fmt.Println("StaTusCode invalid")
	}

}

func bodyContains(Contains interface{}, respObj interface{}) {

	if Contains.(map[string]interface{})["Type"] == "List" {

		switch respObj.(type) {
		case []interface{}:
			fmt.Println("is list true")
		default:
			fmt.Println("is list false")
		}

		list(Contains, respObj)
	} else if Contains.(map[string]interface{})["Type"] == "Object" {
		switch respObj.(type) {
		case map[string]interface{}:
			fmt.Println("is obj true")
		default:
			fmt.Println("is obj false")
		}
		object(Contains, respObj)
	} else {
		fmt.Println("error: no body")
	}

}

func list(Contains interface{}, respObj interface{}) {
	Lenght := Contains.(map[string]interface{})["Lenght"]
	if Lenght.(map[string]interface{})["Equal"] != nil {
		leng := Lenght.(map[string]interface{})["Equal"]
		Equal(leng.(int), len(respObj.([]interface{})))
	}
	InType := Contains.(map[string]interface{})["InType"]
	if InType != nil {
		listValue(InType, respObj)
	}

}

func listValue(InType interface{}, respObj interface{}) {
	intype := InType.([]interface{})[0]
	//fmt.Println(intype)

	for key, value := range intype.(map[interface{}]interface{}) {
		obj := respObj.([]interface{})[key.(int)]
		switch value.(type) {
		case map[string]interface{}:
			fmt.Println(key, reflect.TypeOf(value))
			bodyContains(value.(map[string]interface{})["Contains"], obj)
			//fmt.Println("Integer:", value.(map[string]interface{})["Contains"])
		default:
			fmt.Println(key, reflect.TypeOf(value))
		}

	}
}

func object(Contains interface{}, respObj interface{}) {
	Lenght := Contains.(map[string]interface{})["Lenght"]
	if Lenght.(map[string]interface{})["Equal"] != nil {
		leng := Lenght.(map[string]interface{})["Equal"]
		Equal(leng.(int), len(respObj.(map[string]interface{})))
	}
	InType := Contains.(map[string]interface{})["InType"]

	//fmt.Println(InType)
	if InType != nil {
		objValue(InType, respObj)
	}
}

func objValue(InType interface{}, respObj interface{}) {
	intype := InType.([]interface{})[0]
	//fmt.Println(intype)

	for key, value := range intype.(map[string]interface{}) {
		obj := respObj.(map[string]interface{})[key]
		fmt.Println(key, reflect.TypeOf(value), obj)

		val := value.(map[string]interface{})["Contains"]
		if val != nil {
			bodyContains(value, obj)
		}

	}
}

func Equal[T int | string](yml T, resp T) {
	fmt.Println("equal", yml, resp)
	if yml == resp {
		fmt.Println("true")
	}
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
