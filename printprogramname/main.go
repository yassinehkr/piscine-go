package main

import (
		"github.com/01-edu/z01"
		"os"
)

func main() {
	p := os.Args[0]

	for _, char := range p {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n') 
}
