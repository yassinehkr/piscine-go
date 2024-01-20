package piscine

func AppendRange(min, max int) []int {
	a := []int{}
	for x := min; x < max; x++ {

		a = append(a, x)
	}
	return a

}
