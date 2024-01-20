package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	x := make([]int, max-min)
	for i := min; i < max; i++ {
		x[i-min] = i
	}
	return x
}
