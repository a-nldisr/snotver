package main

import (
	"flag"
	"fmt"
	"os"
)


const appVersion = "2.0"

func main() {

	version := flag.Bool("v", false, "Prints the version of the program")
	flag.Parse()
	if *version {
		fmt.Println("Version: " + appVersion)
		os.Exit(0)
	}
	fmt.Println("Please use the -v flag")
	os.Exit(1)
}
