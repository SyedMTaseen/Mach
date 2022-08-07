package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

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
		request(testcase, RequestURL, HTTPmethods)

	}

	//fmt.Print(RequestURL)

}

func request(testcase interface{}, RequestURL interface{}, HTTPmethods string) {
	Request := testcase.(map[string]interface{})["Request"]
	url := addParams(RequestURL.(string), Request)
	fmt.Println(url)
	Body := getBody(Request)
	fmt.Println(Body)

	responce := RESTRequest(Request, HTTPmethods, url, Body)

	respObj := ResponcetoObject(responce)
	//fmt.Print(RequestURL)
	fmt.Println(responce.StatusCode)
	// vars, err := respObj.(map[string]interface{})["city"]
	// if err == true {
	// 	fmt.Print("error")
	// }
	fmt.Println(reflect.TypeOf(respObj))
}

func addParams(RequestUrl string, Request interface{}) string {

	Params := Request.(map[string]interface{})["Params"]
	if Params == nil {
		return RequestUrl
	}
	parms := Params.([]interface{})[0]
	requestUrl := RequestUrl
	endprams := ""
	for key, value := range parms.(map[string]interface{}) {
		if strings.Contains(requestUrl, "{{"+key+"}}") {
			requestUrl = strings.ReplaceAll(requestUrl, "{{"+key+"}}", value.(string))
		} else {
			endprams += key + "=" + value.(string)
		}
		//fmt.Println(key, value)
	}
	if endprams != "" {
		requestUrl += "?" + endprams
	}
	return requestUrl
}

func getBody(Request interface{}) *bytes.Buffer {
	Body := Request.(map[string]interface{})["Body"]
	if Body == nil {
		return nil
	}
	in := []byte(Body.(string))
	return bytes.NewBuffer(in)
}

func addHeader(Request interface{}, req *http.Request) *http.Request {
	Header := Request.(map[string]interface{})["Header"]
	if Header != nil {
		header := Header.([]interface{})[0]
		for key, value := range header.(map[string]interface{}) {
			req.Header.Add(key, value.(string))
			//fmt.Println(key, value)
		}
	}
	return req
}

func RESTRequest(Request interface{}, HTTPmethods string, url string, Body *bytes.Buffer) *http.Response {
	var req *http.Request
	if Body == nil {
		reqnullbody, _ := http.NewRequest(HTTPmethods, url, nil)
		req = reqnullbody
	} else {
		reqwithbody, _ := http.NewRequest(HTTPmethods, url, Body)
		req = reqwithbody
	}

	req = addHeader(Request, req)
	res, _ := http.DefaultClient.Do(req)

	return res
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
