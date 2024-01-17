package piscine

func IsAlpha(s string) bool {
	bol := true
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z' || letter >= '0' && letter <= '9' {
			bol = true
			continue
		} else {
			bol = false
		}
	}
	return bol
}
