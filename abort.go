package piscine

func Abort(a, b, c, d, e int) int {
	res := 0
	for x := 0; x < 5; x++ {
		res = a + b + c + d + e
	}
	return res / 5
}
