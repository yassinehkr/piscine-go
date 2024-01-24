package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return nil
	}
	a := []int{}
	for x := max; x > min; x-- {
		a = append(a, x)
	}
	return a
}
