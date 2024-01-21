package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	x := os.Args[0]
	for _, v := range x {
		z01.PrintRune(v)
	}
	{
		z01.PrintRune('\n')
	}
}
