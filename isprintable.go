package piscine

func IsPrintable(s string) bool {
	r := []rune(s)
	for x := 0; x < len(s); x++ {
		if r[x] < 31 {
			return false
		}
	}
	return true
}
