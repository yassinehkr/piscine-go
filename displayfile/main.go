package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	s := os.Args

	if len(s) > 2 {
		fmt.Printf("Too many arguments")
	} else if len(s) < 2 {
		fmt.Printf("File name missing")
	} else {
		content, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf(err.Error())
		} else {
			fmt.Printf(string(content))
		}
	}
}
