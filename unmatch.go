package piscine

func Unmatch(arr []int) int {
	for i := 0; i < len(arr); i++ {
		c := 0

		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				c++
			}
		}

		if c%2 != 0 {
			return arr[i]
		}
	}

	return -1
}
