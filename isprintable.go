package piscine

func IsPrintable(s string) bool {
	for _, char := range s {
		if char > 126 || char < 32 {
			return false
		}
	}
	return true
}
