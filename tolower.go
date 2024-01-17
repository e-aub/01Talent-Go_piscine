package piscine

func ToLower(s string) string {
	sRune := []rune(s)
	for index, char := range sRune {
		if char >= 'A' && char <= 'Z' {
			sRune[index] = sRune[index] + 32
		} else {
			continue
		}
	}
	return string(sRune)
}
