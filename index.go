package piscine

func Index(s string, toFind string) int {
	if s == "" {
		return -1
	} else if toFind == "" {
		return 0
	}
	sR := []rune(s)
	toFindRune := []rune(toFind)

	for i := 0; i < len(sR); i++ {
		bol := true
		for j := 0; j < len(toFind); j++ {
			if sR[i+j] != toFindRune[j] {
				bol = false
				break
			}
		}
		if bol {
			return i
		}
	}

	return -1
}
