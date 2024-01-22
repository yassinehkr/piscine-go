package main

import (
	"github.com/01-edu/z01"
)

const a = "x = 42, y = 21"

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	for _, r := range a {
		z01.PrintRune(r)
	}
}
