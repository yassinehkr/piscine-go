package piscine

func ReverseMenuIndex(menu []string) []string {
	name := make([]string, len(menu))

	for a := 0; a < len(menu); a++ {
		name[len(menu)-1-a] = menu[a]
	}
	return name
}
