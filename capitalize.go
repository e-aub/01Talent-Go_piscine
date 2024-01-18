package piscine

func Capitalize(s string) string {
	sRune := []rune(s)
	for index, letter := range sRune {
		if letter >= 'A' && letter <= 'Z' {
			sRune[index] += 32
		}
	}
	if sRune[0] >= 'a' && sRune[0] <= 'z' {
		sRune[0] -= 32
	}
	for i := 0; i < len(sRune)-1; i++ {
		if !(sRune[i] >= 'a' && sRune[i] <= 'z' || sRune[i] >= '0' && sRune[i] <= '9' || sRune[i] >= 'A' && sRune[i] <= 'Z') {
			if sRune[i+1] >= 'a' && sRune[i+1] <= 'z' {
				sRune[i+1] -= 32
			}
		}
	}
	return string(sRune)
}
