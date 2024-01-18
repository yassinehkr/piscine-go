package piscine

import "github.com/01-edu/z01"

func sort(tab []int) {
	l := 0
	for j := range tab {
		l++
		j++
	}
	i := 1
	for lg > i {
		if tab[i-1] > tab[i] {
			tab[i], tab[i-1] = tab[i-1], tab[i]
			i = 1
		} else {
			i++
		}
	}
}

func PrintNbrInOrder(n int) {
	var tab []int
	if n == 0 {
		tab = append(tab, 0)
	}
	for i := 1; n > 0; i++ {
		tab = append(tab, n%10)
		n /= 10
	}
	sort(tab)
	for i := range tab {
		z01.PrintRune(rune(tab[i] + '0'))
	}
}
