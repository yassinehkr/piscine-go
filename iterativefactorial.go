package main

import "fmt"

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}

func IterativeFactorial(nb int) int {
	if nb == 1 {
		return 1
	} else if nb >= 24 {

		return 0
	}
	return nb * IterativeFactorial(nb-1)
}
