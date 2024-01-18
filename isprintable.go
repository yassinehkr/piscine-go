package piscine

func IsPrintable(s string) bool {
	r := []rune(s)
	for x := 0; x < len(s); x++ {
		if r[x] <= 32 || r[x] >= 126 {
			return true
		}
	}
	return false
}
