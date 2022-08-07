package pkg

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func Request(testcase interface{}, RequestURL interface{}, HTTPmethods string) *http.Response {
	Request := testcase.(map[string]interface{})["Request"]
	url := AddParams(RequestURL.(string), Request)
	fmt.Println(url)
	Body := GetBody(Request)
	fmt.Println(Body)

	responce := RESTRequest(Request, HTTPmethods, url, Body)

	return responce
}

func AddParams(RequestUrl string, Request interface{}) string {

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

func GetBody(Request interface{}) *bytes.Buffer {
	Body := Request.(map[string]interface{})["Body"]
	if Body == nil {
		return nil
	}
	in := []byte(Body.(string))
	return bytes.NewBuffer(in)
}

func AddHeader(Request interface{}, req *http.Request) *http.Request {
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

	req = AddHeader(Request, req)
	res, _ := http.DefaultClient.Do(req)

	return res
}
