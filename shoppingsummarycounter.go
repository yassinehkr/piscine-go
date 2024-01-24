package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	start := 0
	item := ""
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' || i == len(str)-1 {
			if i == len(str)-1 {
				item = str[start:]
			} else {
				item = str[start:i]
			}
			summary[item]++
			start = i + 1
		}
	}

	return summary
}
