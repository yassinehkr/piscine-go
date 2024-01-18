package piscine

func IsNumeric(s string) bool {
	r := []rune(s)
	for x := 0; x < len(s); x++ {
		if r[x] < '0' || r[x] > '9' {
			return false
		}
	}
	return true
}
