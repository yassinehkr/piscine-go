package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	s := os.Args

	if len(s) > 2 {
		fmt.Println("Too many arguments")
	} else if len(s) < 2 {
		fmt.Println("File name missing")
	} else {
		content, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(content))
		}
	}
}
