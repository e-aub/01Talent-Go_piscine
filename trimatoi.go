package piscine

func TrimAtoi(s string) int {
	var result int
	sign := 1
	signcount := 0
	digcount := 0

	for _, digit := range s {
		if digit >= '0' && digit <= '9' {
			result = (result * 10) + int(digit-'0')
			digcount++
		}
		if digit == '-' && signcount == 0 && digcount == 0 {
			sign = -1
			signcount++
		}
	}

	return result * sign
}
