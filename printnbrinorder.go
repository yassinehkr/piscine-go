package piscine

import "github.com/01-edu/z01"

func PrintDigitsInOrder(n int) {
	if n == 0 {
		z01.PrintRune("0")
		return
	}

	r := ""
	for n > 0 {
		r = string(rune(n%10)+'0') + r
		n /= 10
	}

	str := []rune(r)
	i := 1
	for i < len(str) {
		if str[i] < str[i-1] {
			str[i], str[i-1] = str[i-1], str[i]
			if i > 1 {
				i--
			}
		} else {
			i++
		}
	}

	z01.PrintRune(string(str))
}
