package piscine

func IterativeFactorial(nb int) int {
	if nb == 1 {
		return 1
	} else if nb >= 15 {

		return 0
	}
	return nb * IterativeFactorial(nb-1)
}
