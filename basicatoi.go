package piscine

func BasicAtoi(s string) int {
	len := len(s) - 1
	num := []rune(s)
	res := 0
	for i := 0; len >= i; i++ {
		if num[i] >= '0' && num[i] <= '9' {
			res = res*10 + int(num[i]-'0')
		} else {
			return 0
		}
	}
	return res
}
