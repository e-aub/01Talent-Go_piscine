package piscine

func BasicAtoi2(s string) int {
	char := []rune(s)
	res := 0
	len := len(s)
	for i := 0; i < len; i++ {
		if char[i] >= '0' && char[i] <= '9' {
			res = res*10 + int(char[i]-'0')
		} else {
			return 0
		}
	}
	return res
}
