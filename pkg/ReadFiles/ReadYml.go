package ReadFiles

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func ReadYml() map[interface{}]interface{} {
	buf := readFile()
	//Map the read file[interface {}]interface {}Map to
	return mapYmltoInterface(buf)
}

func readFile() []byte {
	buf, err := ioutil.ReadFile("/workspaces/Mach/Apitest.yml")
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
