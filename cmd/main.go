package main

import (
	"Mach/pkg/readFiles"
	"flag"
	"fmt"
	"os"
)

func main() {

	//test.read yml

	getCmd := flag.NewFlagSet("test", flag.ExitOnError)
	getpath := getCmd.String("path", "", "testcases path")

	if len(os.Args) < 2 {
		fmt.Println("expected 'test' as subcommands")
		os.Exit(1)
	}

	//look at the 2nd argument's value
	switch os.Args[1] {
	case "test": // if its the 'get' command
		HandleTest(getCmd, getpath)
	case "-version": // if its the 'get' command
		fmt.Println("v0.1.1")
	case "-help": // if its the 'get' command
		fmt.Println("")
		fmt.Println("Usage : mach test -path /path/to/the/testcases/")
		fmt.Println("")
	default:
		fmt.Println("Error not a valid command user -help for more info") // if we don't understand the input
	}

	//performMach(t)

}

func HandleTest(getCmd *flag.FlagSet, path *string) {

	getCmd.Parse(os.Args[2:])
	if *path == "" {
		fmt.Print("Path is required")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	if *path != "" {
		path := *path
		readFiles.ReadYml(path)
		//fmt.Println(path)
	}
}
