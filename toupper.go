package piscine

func ToUpper(s string) string {
	sRune := []rune(s)
	for index, letter := range sRune {
		if letter >= 'A' && letter <= 'Z' {
			continue
		} else if letter >= 'a' && letter <= 'z' {
			sRune[index] = sRune[index] - 32
		}
	}
	return string(sRune)
}
