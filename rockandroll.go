package piscine

func RockAndRoll(n int) string {
	two := "rock\n"
	three := "roll\n"
	twoThree := "rock and roll\n"
	negative := "error: number is negative\n"
	nonDivisible := "error: non divisible\n"

	if n < 0 {
		return negative
	} else if n == 0 {
		return nonDivisible
	} else if n%2 == 0 && n%3 == 0 {
		return twoThree
	} else if n%2 == 0 {
		return two
	} else if n%3 == 0 {
		return three
	}
	return nonDivisible
}
