package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args[1:]
	for _, v := range a {
		if v == "01" || v == "galaxy" || v == "galaxy 01" {
			fmt.Println("Alert!!!")
		}
	}
}
