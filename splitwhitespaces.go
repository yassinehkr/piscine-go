package piscine

func SplitWhiteSpaces(s string) []string {
	r := []string{}
	b := ""

	for _, char := range s {
		if char == ' ' || char == '\n' || char == '\t' {
			if b != "" {
				r = append(r, b)
				b = ""
			}
		} else {
			b += string(char)
		}
	}

	if b != "" {
		r = append(r, b)
	}

	return r
}
