package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 31 {
		return 0
	} else {
		fact := 1
		for i := 1; i <= nb; i++ {
			fact = fact * i
		}
		return fact
	}
}
