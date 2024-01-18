package piscine

func ToUpper(s string) string {
	r := []rune(s)
	for x, y := range r {
		if y >= 'a' && y <= 'z' {
			r[x] = y - ('a' - 'A')
		}
	}
	return string(r)
}
