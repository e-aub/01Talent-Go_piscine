package piscine

func Capitalize(s string) string {
	sRune := []rune(s)
	for index, char := range sRune {
		if char >= 'A' && char <= 'Z' {
			sRune[index] = sRune[index] + 32
		} else {
			continue
		}
	}

	for index, letter := range sRune {
		if (letter >= 'a' && letter <= 'z') && index == 0 {
			sRune[index] = sRune[index] - 32
		} else if (letter >= 'a' && letter <= 'z') && (sRune[index-1] <= 44 && sRune[index-1] >= 32) || (sRune[index-1] <= 63 && sRune[index-1] >= 58) {
			sRune[index] = sRune[index] - 32
		}
	}
	return string(sRune)
}
