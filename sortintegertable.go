package piscine

func SortIntegerTable(table []int) {
	l := len(table)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if table[j] > table[j+1] {
				table[j+1], table[j] = table[j], table[j+1]
			}
		}
	}
}
