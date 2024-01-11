package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	char := []rune(s)
	res := 0
	sign := 1

	if char[0] == '-' {
		char = char[1:]
		sign = -1
	}

	if char[0] == '+' {
		char = char[1:]
	}
	len := len(char)
	for i := 0; i < len; i++ {
		if char[i] >= '0' && char[i] <= '9' {
			res = res*10 + int(char[i]-'0')
		} else {
			return 0
		}
	}
	return res * sign
}
