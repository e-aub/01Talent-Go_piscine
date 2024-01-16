package piscine

func NRune(s string, n int) rune {
	nrune := []rune(s)
	if n < 0 || n > len(s) {
		return 0
	} else if n == 0 {
		return nrune[0]
	} else {
		return nrune[n-1]
	}
}
