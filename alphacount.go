package piscine

func AlphaCount(s string) int {
	r := []rune(s)
	x := 0
	for i := 0; i < len(r); i++ {
		if (r[i] >= 'a' && r[i] <= 'z' || r[i] >= 'A' && r[i] <= 'Z') {
			x++

		}

	}
	return x
}
