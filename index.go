package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	for x := 0; x <= len(s)-len(toFind); x++ {
		y := s[x : x+len(toFind)]
		if toFind == y {
			return x
		}
	}
	return -1
}
