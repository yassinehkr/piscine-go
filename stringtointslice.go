package main

import "fmt"

func StringToIntSlice(str string) []int {
	var result []int
	for _, v := range str {
		result = append(result, int(v))
	}
	return result
}

func main() {
	fmt.Println(StringToIntSlice("A quick brown fox jumps over the lazy dog"))
	fmt.Println(StringToIntSlice("Converted this string into an int"))
	fmt.Println(StringToIntSlice("hello THERE"))
}
