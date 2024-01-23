package piscine

func Any(f func(string) bool, a []string) bool {
	for _, letter := range a {
		if f(letter) {
			return true
		}
	}
	return false
}
