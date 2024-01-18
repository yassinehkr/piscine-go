package piscine

func IsAlpha(s string) bool {
	r := []rune(s)
	for x := 0; x < len(s); x++ {
		if r[x] < 'a' || r[x] > 'z' {
			if r[x] < 'A' || r[x] > 'Z' {
				if r[x] < '0' || r[x] > '9' {
					return false
				}
			}
		}
	}
	return true
}
