package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	a := []int{}
	for x := min; x < max; x++ {

		a = append(a, x)
	}
	return a

}
