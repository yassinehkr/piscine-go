package piscine

func IterativeFactorial(nb int) int {
	if nb > 1 && nb < 2147483647 {
		r := 1
		for i := nb; i >= 1; i-- {
			r *= i
		}
		return int(r)
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		return 0
	}
}
