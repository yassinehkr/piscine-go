package piscine

func AlphaCount(s string) int {
	r := []rune(s)
	a := 0
	for x := 0; x < len(s); x++ {
		if (r[x] >= 'a' && r[x] <= 'z') || (r[x] >= 'A' && r[x] <= 'Z') {
			a++
		}
	}
	return a
}
