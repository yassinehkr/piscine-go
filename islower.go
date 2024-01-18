package piscine

func IsUpper(s string) bool {
	r := []rune(s)
	for x := 0; x < len(s); x++ {
		if r[x] < 'a' || r[x] > 'z' {
			return false
		}
	}
	return true
}
