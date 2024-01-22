package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	p := os.Args
	name := p[0]

	r := []rune(name)
	for i := range r {
		z01.PrintRune(r[i])
	}
	z01.PrintRune('\n')
}
