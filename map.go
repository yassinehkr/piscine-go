package piscine

func Map(f func(int) bool, a []int) []bool {
	x := make([]bool, len(a))
	for i := range a {
		x[i] = f(a[i])
	}
	return x
}
