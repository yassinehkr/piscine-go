package piscine

func MakeRange(min, max int) []int {
	if max < min {
		return nil
	}
	a := make([]int, max-min)
	for i := min; i < max; i++ {
		a[i-min] = i
	}
	return a
}
