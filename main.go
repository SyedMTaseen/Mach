package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/yaml.v3"
)

func main() {
	// test.read yml
	buf, err := ioutil.ReadFile("Apitest.yml")
	if err != nil {
		fmt.Print("error: Failed to read the file\n")
		return
	}

	//Map the read file[interface {}]interface {}Map to
	t := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &t)
	if err != nil {
		panic(err)
	}

	fmt.Print(t["APITESTING"].(map[string]interface{})["Name"]) // Name
	fmt.Print("\n")
	RequestURL := t["APITESTING"].(map[string]interface{})["RequestURL"]
	fmt.Print(RequestURL)
	fmt.Print("\n")
	//fmt.Print(t["APITESTING"].(map[string]interface{})["TestCases"].([]interface{})[0])

	testcases := t["APITESTING"].(map[string]interface{})["TestCases"].([]interface{})
	testcasesLen := len(testcases)
	fmt.Print(testcasesLen)
	fmt.Print("\n")

	for i := 0; i < testcasesLen; i++ {
		testcase := testcases[i]
		Name := testcase.(map[string]interface{})["Name"]
		fmt.Println(Name)
		Request := testcase.(map[string]interface{})["Request"]

		Params := Request.(map[string]interface{})["Params"]

		parms := Params.([]interface{})[0]
		for key, value := range parms.(map[string]interface{}) {
			fmt.Println(key, value)
		}

		Body := Request.(map[string]interface{})["Body"]
		fmt.Println(Body)

		Header := Request.(map[string]interface{})["Header"]

		if Header != nil {
			header := Header.([]interface{})[0]
			for key, value := range header.(map[string]interface{}) {
				fmt.Println(key, value)
			}
		}

		url := "https://test.api/api/users"

		payload := strings.NewReader("name=test&jab=teacher")

		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("content-type", "application/x-www-form-urlencoded")
		req.Header.Add("cache-control", "no-cache")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(string(body))
		//reader := strings.NewReader(`{"body":123}`)
		// curl := exec.Command("curl", "-l", "-X", "https://api.test.io/?name=bella")
		// output, e := curl.Output()
		// fmt.Println(output)
		// fmt.Println(e)

	}
	//fmt.Print(testcase.(map[string]interface{})["Name"])
	// t["secound"]To(map[interface {}]interface {})Type conversion with
	// fmt.Print(t["secound"].(map[string]interface{})["a1"]) //count1
	// fmt.Print("\n")

	// len()Returns the number of elements in the array
	// 	fmt.Print(len(t["secound"].(map[string]interface{})))
	// 	fmt.Print("\n")

	// 	// []interface {}Array of types
	// 	fmt.Print(t["secound"].(map[string]interface{})["a2"].([]interface{})[0]) // count2
	// 	fmt.Print("\n")

	// 	fmt.Print(t["secound"].(map[string]interface{})["a3"].(map[string]interface{})["b1"]) // count4
	// 	fmt.Print("\n")

	// 	//If it is named regularly, you can check how many there are
	// 	flag, i := 0, 0
	// 	for flag == 0 {
	// 		i++
	// 		switch t["secound"].(map[string]interface{})["a"+strconv.Itoa(i)].(type) {
	// 		case nil:
	// 			flag = 1
	// 		}
	// 	}
	// 	fmt.Printf("a%d is not found\n", i) // a4 is not found

}
