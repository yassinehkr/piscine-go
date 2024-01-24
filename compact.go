package piscine

func Compact(ptr *[]string) int {
	r := 0
	for _, s := range *ptr {
		if s != "" {
			r++
		}
	}
	count := 0
	array := make([]string, r)
	for _, s := range *ptr {
		if s != "" {
			array[count] = s
			count++
		}
	}
	*ptr = array
	return r
}
