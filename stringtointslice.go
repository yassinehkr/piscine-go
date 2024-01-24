package piscine

func StringToIntSlice(str string) []int {
	result := []int{}
	for _, v := range str {
		result = append(result, int(v))
	}
	return result
}
