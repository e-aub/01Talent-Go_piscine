package piscine

func Compare(a, b string) int {
	aRune := []rune(a)
	bRune := []rune(b)
	var length int
	if len(aRune) > len(bRune) {
		length = len(bRune)
	} else if len(aRune) < len(bRune) {
		length = len(aRune)
	} else if len(aRune) == len(bRune) {
		length = len(aRune)
	}

	for i := 0; i < length; i++ {
		if aRune[i] < bRune[i] || len(aRune) < len(bRune) {
			return -1
		} else if aRune[i] > bRune[i] || len(aRune) > len(bRune) {
			{
				return 1
			}
		}
	}
	return 0
}
