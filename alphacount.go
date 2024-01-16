package piscine

func AlphaCount(s string) int {
	count := 0
	for _, letter := range s {
		if letter >= 'A' && letter <= 'Z' || letter >= 'a' && letter <= 'z' {
			count++
		} else {
			continue
		}
	}
	return count
}
