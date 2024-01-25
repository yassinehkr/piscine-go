package piscine

func ActiveBits(n int) int {
	count := 0
	if n == 0 {
		return 0
	}
	for n > 0 {
		if n%2 == 0 {
			n = n / 2
		} else if n%2 != 0 {
			n = n / 2
			count++
		}
	}
	return count
}
