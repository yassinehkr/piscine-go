package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	r := 1
	for x := 0; x < power; x++ {
		r *= nb
	}
	{
		return r
	}
}
