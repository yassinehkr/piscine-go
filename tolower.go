package piscine

func ToLower(s string) string {
	r := []rune(s)
	for x, y := range r {
		if y >= 'A' && y <= 'Z' {
			r[x] = y - ('A' - 'a')
		}
	}
	return string(r)
}
