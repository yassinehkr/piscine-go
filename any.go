package piscine

func Any(f func(string) bool, a []string) bool {
	r := false
	for _, str := range a {
		r := f(str)
		if r == true {
			return r
		} else {
			continue
		}
	}
	return r
}
