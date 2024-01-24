package piscine

func JumpOver(str string) string {
	var s string
	for i := 2; i < len(str); i++ {
		s = s + string(str[i])
		i += 2
	}
	return s
}
