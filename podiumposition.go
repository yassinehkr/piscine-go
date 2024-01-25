package piscine

func PodiumPosition(podium [][]string) [][]string {
	podium[0], podium[1], podium[2], podium[3] = podium[3], podium[2], podium[1], podium[0]
	return podium
}
