package piscine

func Index(s string, toFind string) int {
	toFindRune := []rune(toFind)
	for index, letter := range s {
		if letter == toFindRune[0] {
			return index
		}
	}
	return -1
}
