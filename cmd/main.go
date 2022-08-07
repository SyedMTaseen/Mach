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
		respObj := ResponcetoObject(responce)
		//fmt.Print(RequestURL)
		fmt.Println(responce.StatusCode)
		// vars, err := respObj.(map[string]interface{})["city"]
		// if err == true {
		// 	fmt.Print("error")
		// }
		fmt.Println(respObj)
		//ResponceMain(testcase)
		//
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
		object(Contains)
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
		listValue(InType)
	}

}

func listValue(InType interface{}) {
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

func object(Contains interface{}) {
	Lenght := Contains.(map[string]interface{})["Lenght"]
	if Lenght.(map[string]interface{})["Equal"] != nil {
		Equal(Lenght)
	}
	InType := Contains.(map[string]interface{})["InType"]

	//fmt.Println(InType)
	if InType != nil {
		objValue(InType)
	}
}

func objValue(InType interface{}) {
	intype := InType.([]interface{})[0]
	//fmt.Println(intype)

	for key, value := range intype.(map[string]interface{}) {

		fmt.Println(key, reflect.TypeOf(value))

		val := value.(map[string]interface{})["Contains"]
		if val != nil {
			bodyContains(value)
		}

	}
}

func Equal(Lenght interface{}) {
	fmt.Println(Lenght)
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
