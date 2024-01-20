package piscine

func ConcatParams(args []string) string {
	x := 0
	str := ""
	for range args {
		x++
	}
	for i, v := range args {
		str += v
		if i != x-1 {
			str += "\n"
		}
	}
	return str
}
