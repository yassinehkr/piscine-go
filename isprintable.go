package piscine

func IsPrintable(s string) bool {
	r := []rune(s)
	for x := 0; x < len(s); x++ {
		if r[x] < '0' {
			return true
		}
	}
	return false
}
