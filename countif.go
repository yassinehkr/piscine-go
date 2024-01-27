package piscine

func CountIf(f func(string) bool, tab []string) int {
	c := 0
	for _, s := range tab {
		if f(s) == true {
			c++
		}
	}
	return c
}
