package ReadFiles

import (
	"Mach/pkg/Logger"
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
		Logger.ErrorLogger.Println("error: Failed to read the file\n" + ":" + err.Error())
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
