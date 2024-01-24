package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		found := false
		for j := 0; j < len(a); j++ {
			if i != j && a[i] == a[j] {
				found = true
				break
			}
		}
		if !found {
			return a[i]
		}
	}
	return -1
}
