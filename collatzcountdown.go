package piscine

func CollatzCountdown(start int) int {
	nb := start
	steps := 0

	if start <= 0 {
		return -1
	}

	for nb != 1 {
		if nb%2 == 0 {
			nb = nb / 2
		} else {
			nb = nb*3 + 1
		}
		steps++
	}

	return steps
}
