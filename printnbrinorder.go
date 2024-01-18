package piscine

func PrintDigitsInOrder(n int) {
	if n == 0 {
		fmt.Print("0")
		return
	}

	r := ""
	for n > 0 {
		r = string(rune(n%10)+'0') + r
		n /= 10
	}

	s := []rune(r)
	i := 1
	for i < len(s) {
		if s[i] < s[i-1] {
			s[i], s[i-1] = s[i-1], s[i]
			if i > 1 {
				i--
			}
		} else {
			i++
		}
	}

	fmt.Print(string(s))
}
