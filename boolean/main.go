package main

import (
	"fmt"
	"os"
)

const (
	EvenMsg = "I have an even number of arguments"
	OddMsg  = "I have an odd number of arguments"
)

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	lengthOfArg := len(os.Args[1:])
	if isEven(lengthOfArg) {
		fmt.Println(EvenMsg)
	} else {
		fmt.Println(OddMsg)
	}
}
