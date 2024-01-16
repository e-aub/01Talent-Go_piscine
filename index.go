package piscine

func Index(s string, toFind string) int {
	if s == "" || toFind == "" {
		return -1
	}
	toFindRune := []rune(toFind)
	for index, letter := range s {
		if letter == toFindRune[0] {
			return index
		}
	}
	return -1
}
