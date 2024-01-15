package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else {
		fact := 1
		for i := 1; i <= nb; i++ {
			fact = fact * i
			if fact >= 9223372036854775807 {
				return 0
			}
		}
		return fact
	}
}
