package piscine

func Abort(a, b, c, d, e int) int {
	tab := [5]int{a, b, c, d, e}
	l := 0
	for j := range tab {
		l++
		j++
	}
	i := 1
	for l > i {
		if tab[i-1] > tab[i] {
			tab[i], tab[i-1] = tab[i-1], tab[i]
			i = 1
		} else {
			i++
		}
	}
	return tab[2]
}
